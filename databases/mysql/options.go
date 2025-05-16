package mysql

import (
	"errors"
	"fmt"
	"time"

	"github.com/gscaffold/helpers/databases"
	"github.com/gscaffold/helpers/devops"
	"github.com/gscaffold/helpers/internal/cgroup"
	"github.com/gscaffold/utils"
	"gorm.io/gorm/logger"
)

type Options struct {
	master string
	slaves []string

	logger                 logger.Interface
	logLevel               logger.LogLevel
	skipDefaultTransaction bool // 默认跳过单条事务, 提升 30% 性能.
	maxIdleConns           int
	maxOpenConns           int
	maxLifeTime            time.Duration
}

func (opts *Options) Validate() error {
	if opts.master == "" {
		return errors.New("must have master dsn")
	}
	return nil
}

func initOptions(app, name string) (*Options, error) {
	master, err := devops.Discovery(devops.ResourceMySQL, app, name, "master")
	if err != nil {
		return &Options{}, err
	}

	slaves, err := devops.DiscoveryMany(devops.ResourceMySQL, app, name, "slave")
	if err != nil {
		return &Options{}, err
	}

	logLevel := logger.Info
	if utils.IsProd() {
		logLevel = logger.Warn
	}

	return &Options{
		master:                 master,
		slaves:                 slaves,
		logLevel:               logLevel,
		skipDefaultTransaction: true,
		maxIdleConns:           cgroup.TotalCPU(),
		maxOpenConns:           30 * cgroup.TotalCPU(),
		maxLifeTime:            time.Minute * 5,
	}, nil
}

type Option func(opts *Options)

func OptionOpenConfig(cfg databases.OpenConfig) Option {
	return func(opts *Options) {
		if cfg.DSN != "" {
			opts.master = cfg.DSN
			return
		}
		opts.master = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.User, cfg.Password, cfg.Address, cfg.Name)
	}
}

func OptionSlavesOpenConfig(cfgs []databases.OpenConfig) Option {
	return func(opts *Options) {
		for _, cfg := range cfgs {
			if cfg.DSN != "" {
				opts.slaves = append(opts.slaves, cfg.DSN)
				continue
			}
			dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				cfg.User, cfg.Password, cfg.Address, cfg.Name)
			opts.slaves = append(opts.slaves, dsn)
		}
	}
}

func OptionLogger(logger logger.Interface) Option {
	return func(opts *Options) {
		opts.logger = logger
	}
}

// WithLoggerLevel 和 WithLogger 互斥, 优先采用 WithLogger
func OptionLoggerLevel(level logger.LogLevel) Option {
	return func(opts *Options) {
		opts.logLevel = level
	}
}

func OptionDefaultTransaction() Option {
	return func(opts *Options) {
		opts.skipDefaultTransaction = false
	}
}
