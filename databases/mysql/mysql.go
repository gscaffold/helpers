package mysql

import (
	"sync/atomic"

	"gorm.io/driver/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
	slaves []*gorm.DB
	next   uint64 // 用于多个读节点之间的负载均衡
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

	slaves := make([]*gorm.DB, 0, len(opts.slaves))
	for _, slaveDSN := range opts.slaves {
		slave, err := open(slaveDSN, opts)
		if err != nil {
			return &DB{}, errors.Wrapf(err, "open slave mysql err, dsn:%s", slaveDSN)
		}
		slaves = append(slaves, slave)
	}

	return &DB{
		DB:     master,
		slaves: slaves,
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

func (db *DB) GetMaster() *gorm.DB {
	return db.DB
}

func (db *DB) GetSlave() *gorm.DB {
	slaveNum := uint64(len(db.slaves))
	if slaveNum == 0 {
		return db.DB
	}
	return db.slaves[atomic.AddUint64(&db.next, 1)%slaveNum]
}

func (db *DB) Close() error {
	err := closeGormDB(db.DB)
	if err != nil {
		return err
	}
	for _, slave := range db.slaves {
		err := closeGormDB(slave)
		if err != nil {
			return err
		}
	}
	return nil
}

func closeGormDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return errors.Wrap(err, "close gorm db error")
	}
	err = sqlDB.Close()
	return errors.Wrap(err, "close gorm db error")
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
