package mysql

import (
	"gorm.io/driver/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

type DB struct {
	*gorm.DB
}

// Discovery 创建一个数据库连接.
// name: 资源名称, 用于资源发现.
func Discovery(name string, _opts ...Option) (*DB, error) {
	return DiscoveryAppExclusive("", name, _opts...)
}

func DiscoveryAppExclusive(app, name string, _opts ...Option) (*DB, error) {
	opts, err := initOptions(app, name, _opts...)
	if err != nil {
		return &DB{}, err
	}
	if err := opts.Validate(); err != nil {
		return &DB{}, err
	}

	if opts.logger == nil {
		opts.logger = logger.Default.LogMode(opts.logLevel)
	}

	master, err := open(opts.master, opts)
	if err != nil {
		return &DB{}, errors.Wrapf(err, "open mysql err, dsn:%s", opts.master)
	}
	var replicas = make([]gorm.Dialector, 0, len(opts.slaves))
	for _, slaveDSN := range opts.slaves {
		replicas = append(replicas, mysql.Open(slaveDSN))
	}
	master.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(opts.master)},
		Replicas: replicas,
		Policy:   dbresolver.RoundRobinPolicy(),
	}))

	return &DB{
		DB: master,
	}, nil

}

// MustDiscovery 是 Discovery 的一个变体，如果出错会 panic
func MustDiscovery(name string, opts ...Option) *DB {
	db, err := Discovery(name, opts...)
	if err != nil {
		panic(err)
	}
	return db
}

func (db *DB) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return errors.Wrap(err, "close gorm db error")
	}
	err = sqlDB.Close()
	if err != nil {
		return errors.Wrap(err, "close gorm db error")
	}
	return nil
}

func open(dsn string, opts *Options) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 opts.logger,
		SkipDefaultTransaction: opts.skipDefaultTransaction,
	})
	if err != nil {
		return &gorm.DB{}, errors.Wrapf(err, "open mysql err, dsn:%s", opts.master)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return &gorm.DB{}, errors.Wrapf(err, "open mysql err, dsn:%s", opts.master)
	}
	sqlDB.SetMaxIdleConns(opts.maxIdleConns)
	sqlDB.SetMaxOpenConns(opts.maxOpenConns)
	sqlDB.SetConnMaxLifetime(opts.maxLifeTime)

	return db, nil
}
