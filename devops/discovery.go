/*
devops 资源发现、服务发现相关代码.

example dsn:
- mysql: user:password@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local
- redis: redis://user:password@localhost:6789/3?dial_timeout=3&db=1&read_timeout=6s&max_retries=2
*/
package devops

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// 注册资源
func Register(resource ResourceType, app, name, kind string, dsns []string) (err error) {
	if resource == "" {
		return errors.New("must have resource")
	}
	if name == "" {
		return errors.New("must have name")
	}
	if app != "" {
		name = app + "__" + name
	}

	prefix := fmt.Sprintf("RES_%s_%s", resource.Upper(), name)
	for i, dsn := range dsns {
		var key string
		if kind == "" {
			key = fmt.Sprintf("%s_%d_ADDR", prefix, i)
		} else {
			key = fmt.Sprintf("%s_%s_%d_ADDR", prefix, kind, i)
		}

		err = os.Setenv(key, dsn)
		if err != nil {
			return err
		}
	}
	return nil
}

// Discovery 资源发现, 通过读取环境变量实现. app、name、kind 中不允许有下划线 _
// name: 资源名称
// kind: 资源的类型, 如 master, slave, 是 name 的进一步细分.
func Discovery(resource ResourceType, app, name, kind string) (string, error) {
	dsns, err := DiscoveryMany(resource, app, name, kind)
	if err != nil {
		return "", err
	}
	if len(dsns) == 0 {
		return "", errors.New("no dsn")
	}
	return dsns[0], nil
}

// DiscoveryMany 资源发现
// app: 默认为空, app 独享资源时使用.
// name: 资源名称
// kind: 资源的类型, 如 master, slave, 是 name 的进一步细分.
func DiscoveryMany(resource ResourceType, app, name, kind string) (dsns []string, _ error) {
	if resource == "" {
		return []string{}, errors.New("must have resource")
	}
	if name == "" {
		return []string{}, errors.New("must have name")
	}
	if app != "" {
		name = app + "__" + name
	}

	prefix := fmt.Sprintf("RES_%s_%s", resource.Upper(), name)
	for i := 0; ; i++ {
		var key string
		if kind == "" {
			key = fmt.Sprintf("%s_%d_ADDR", prefix, i)
		} else {
			key = fmt.Sprintf("%s_%s_%d_ADDR", prefix, kind, i)
		}
		dsn := os.Getenv(key)
		if dsn == "" {
			break
		}
		dsns = append(dsns, dsn)
	}

	return dsns, nil
}

type ResourceType string

func (d ResourceType) Lower() string {
	return strings.ToLower(d.String())
}

func (d ResourceType) Upper() string {
	return strings.ToUpper(d.String())
}

func (d ResourceType) String() string {
	return string(d)
}

const (
	ResourceMySQL  ResourceType = "MySQL"
	ResourceRedis  ResourceType = "Redis" // redis://<user>:<pass>@localhost:6379/<db>
	ResourceKafka  ResourceType = "Kafka"
	ResourceStatsd ResourceType = "Statsd"
)
