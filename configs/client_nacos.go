package configs

import (
	"context"
	"os"

	"github.com/gscaffold/helpers/logger"
	"github.com/gscaffold/utils"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/pkg/errors"
)

type NacosClient struct {
	client config_client.IConfigClient
}

var _ IClient = (*NacosClient)(nil)

const (
	EnvNacosAddr = "env_nacos_addr"
)

func NewNacosClient() (*NacosClient, error) {
	var client config_client.IConfigClient
	{
		addrs := os.Getenv(EnvNacosAddr)
		if addrs == "" {
			utils.HandleFatalError(errors.New("nacos address is empty"),
				"configs", "env_nacos_addr not found")
		}
		sc := []constant.ServerConfig{
			*constant.NewServerConfig(addrs, 8848, constant.WithContextPath("/nacos")),
		}
		logOpt := constant.WithLogLevel("debug")
		if utils.IsProd() {
			logOpt = constant.WithLogLevel("info")
		}
		cc := *constant.NewClientConfig(
			constant.WithNamespaceId(""),
			constant.WithTimeoutMs(5000),
			constant.WithNotLoadCacheAtStart(true),
			constant.WithLogDir("/tmp/nacos/log"),
			constant.WithCacheDir("/tmp/nacos/cache"),
			logOpt,
		)
		var err error
		client, err = clients.NewConfigClient(
			vo.NacosClientParam{
				ClientConfig:  &cc,
				ServerConfigs: sc,
			},
		)
		if err != nil {
			return nil, errors.Wrapf(err, "create nacos client error. config:%+v", cc)
		}
	}
	return &NacosClient{
		client: client,
	}, nil
}

func (client *NacosClient) Name() string {
	return "nacos"
}

func (client *NacosClient) Get(kind, key string) string {
	return client.get(key, kind)
}

func (client *NacosClient) get(kind, key string) string {
	//nolint:exhaustruct
	value, err := client.client.GetConfig(vo.ConfigParam{
		DataId: kind,
		Group:  key,
	})
	if err != nil {
		logger.Errorf(context.TODO(), "nacos get config error. key:%s, kind:%s", key, kind)
		return ""
	}
	return value
}

func (client *NacosClient) BatchGet(kind string, keys ...string) map[string]string {
	return client.batchGet(kind, keys...)
}

func (client *NacosClient) batchGet(kind string, keys ...string) map[string]string {
	result := make(map[string]string, len(keys))
	for _, key := range keys {
		result[key] = client.get(key, kind)
	}
	return result
}

func (client *NacosClient) MonitorChange(kind, key string, fn ConfigChangeCallback) {
	if kind == "" {
		kind = "DEFAULT_GROUP"
	}
	//nolint:exhaustruct
	err := client.client.ListenConfig(vo.ConfigParam{
		Group:  kind,
		DataId: key,
		OnChange: func(_, _, _, data string) {
			fn(data)
		},
	})
	if err != nil {
		logger.Errorf(context.TODO(),
			"nacos listen config error. key:%s, kind:%s, err:%s", key, kind, err.Error())
	}
}
