package redis

import "github.com/redis/go-redis/v9"

type UniversalClient = redis.UniversalClient

func Discovery(name string, _opts ...Option) (UniversalClient, error) {
	return DiscoveryAppExclusive("", name, _opts...)
}

func DiscoveryAppExclusive(app, name string, _opts ...Option) (UniversalClient, error) {
	opts, err := initOptions(app, name, _opts...)
	if err != nil {
		return nil, err
	}
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	client := redis.NewUniversalClient(&opts.UniversalOptions)
	return client, nil
}

func MustDiscovery(name string, _opts ...Option) UniversalClient {
	client, err := DiscoveryAppExclusive("", name, _opts...)
	if err != nil {
		panic(err)
	}
	return client
}
