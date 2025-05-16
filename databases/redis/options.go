package redis

import (
	"time"

	"github.com/gscaffold/helpers/databases"
	"github.com/gscaffold/helpers/devops"
	"github.com/gscaffold/helpers/internal/cgroup"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type Options struct {
	redis.UniversalOptions
}

func (opts *Options) Validate() error {
	if len(opts.Addrs) == 0 {
		return errors.New("must have master dsn")
	}
	return nil
}

func initOptions(app, name string, _opts ...Option) (*Options, error) {
	dsns, err := devops.DiscoveryMany(devops.ResourceRedis, app, name, "")
	if err != nil {
		return &Options{}, err
	}
	var username, password string
	var db int
	addrs := make([]string, 0, len(dsns))
	for _, dsn := range dsns {
		cfg, err := redis.ParseURL(dsn)
		if err != nil {
			return &Options{}, errors.Wrapf(err, "parse redis dsn err, dsn:%s", dsn)
		}
		if len(addrs) == 0 {
			username = cfg.Username
			password = cfg.Password
			db = cfg.DB
		}
		addrs = append(addrs, cfg.Addr)
	}

	opts := &Options{
		UniversalOptions: redis.UniversalOptions{
			Addrs:           addrs,
			Username:        username,
			Password:        password,
			DB:              db,
			MinIdleConns:    cgroup.TotalCPU(),
			MaxIdleConns:    60 * cgroup.TotalCPU(),
			ConnMaxIdleTime: time.Minute * 30,
		},
	}
	for _, opt := range _opts {
		opt(opts)
	}

	return opts, nil
}

type Option func(opts *Options)

func OptionOpenConfig(cfg *databases.OpenConfig) Option {
	return func(opts *Options) {
		if cfg.DSN != "" {
			rcfg, err := redis.ParseURL(cfg.DSN)
			if err != nil {
				return
			}
			opts.Addrs = []string{rcfg.Addr}
			opts.Username = rcfg.Username
			opts.Password = rcfg.Password
			opts.DB = rcfg.DB
		}

		opts.Addrs = []string{cfg.Address}
		opts.Username = cfg.User
		opts.Password = cfg.Password
		opts.DB = cfg.DB
	}
}

func OptionMinIdleConns(cnt int) Option {
	return func(opts *Options) {
		opts.MinIdleConns = cnt
	}
}

func OptionMaxIdleConns(cnt int) Option {
	return func(opts *Options) {
		opts.MaxIdleConns = cnt
	}
}

func OptionConnMaxIdleTime(ttl time.Duration) Option {
	return func(opts *Options) {
		opts.ConnMaxIdleTime = ttl
	}
}
