package sentry

/*
SampleRate 采样
BeforeSend 过滤或定制要发送的事件
sentry.Recover() recover + 上报事件
sentry.Flush()
设置 level, 并且和 logger 交互

资源发现(环境变量)

add logger hook
recover hook: panic 时上报 sentry, 避免程序崩溃, 作为 logger、metrics 的补充.
一般服务场景不存在panic需要重启程序的场景, 如有避免引入该组件.
	add gin hook
	add grpc hook

*/

func Init() {

}
