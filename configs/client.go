package configs

import (
	"sync"

	"github.com/gscaffold/utils"
)

var (
	// 通过更改函数指针实现更好的性能
	getClient func() IClient

	defaultClient IClient
	lock          sync.Mutex
)

func init() {
	getClient = initClientOnce
}

func initClientOnce() IClient {
	lock.Lock()
	defer lock.Unlock()
	if defaultClient == nil {
		var err error
		defaultClient, err = NewNacosClient()
		utils.HandleFatalError(err, "configs", "init source nacos error")

		// 重写函数指针
		getClient = func() IClient {
			return defaultClient
		}
	}
	return defaultClient
}

//go:generate mockgen -destination mocks/client.go -source=client.go
type IClient interface {
	Name() string
	Get(kind, key string) string
	BatchGet(kind string, keys ...string) map[string]string
	MonitorChange(kind, key string, fn ConfigChangeCallback)
}

type ConfigChangeCallback func(newvalue string)
