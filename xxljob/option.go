package xxljob

import (
	"github.com/go-basic/ipv4"
	"github.com/gscaffold/utils"
	xxl "github.com/xxl-job/xxl-job-executor-go"
)

type (
	options struct {
		serverAddr  string
		accessToken string
		executorIP  string
		registryKey string
		logger      xxl.Logger
		port        int
	}
	Option func(*options)
)

func (opts *options) LoadDefault() {
	// if opts.serverAddr == "" {
	// 	opts.serverAddr = ""
	// }
	if opts.accessToken == "" {
		opts.accessToken = "default_token"
	}
	if opts.executorIP == "" {
		opts.executorIP = ipv4.LocalIP()
	}
	if opts.port == 0 {
		opts.port = 9999
	}
	if opts.registryKey == "" {
		opts.registryKey = utils.GetApp()
	}
	if opts.logger == nil {
		opts.logger = &Logger{}
	}
}

func OptionServerAddr(addr string) Option {
	return func(o *options) {
		o.serverAddr = addr
	}
}

// OptionAccessToken 请求令牌
func OptionAccessToken(token string) Option {
	return func(o *options) {
		o.accessToken = token
	}
}

// OptionExecutorIP 设置执行器IP
func OptionExecutorIP(ip string) Option {
	return func(o *options) {
		o.executorIP = ip
	}
}

// OptionExecutorPort 设置执行器端口
func OptionExecutorPort(port int) Option {
	return func(o *options) {
		o.port = port
	}
}

// OptionRegistryKey 设置执行器标识
func OptionRegistryKey(registryKey string) Option {
	return func(o *options) {
		o.registryKey = registryKey
	}
}
