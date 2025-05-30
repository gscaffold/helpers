package trace

/*
jaeger opentracing.SetGlobalTracer

trace 要注意效率, 高QPS接口使用 trace 会影响性能
默认自适应采样, 支持不同采样级别.
ebpf 零代码trace, 性能影响较高(约100ms?)

定义不同的 span
httpSpan, grpcSpan, producerSpan, consumerSpan
mysqlSpan, redisSpan

Spans, 所有span的记录
1. gin
2. grpc
3. P1 kafka consumer?
4. P1 gorm
4. P1 kafka producer?
5. P1 redis?
6. P2: http.DefaultTransport? 提供一个 wrap 方法显式使用, 不要隐式. 也支持Init替换默认实例.
数据库要不要同一个通用的 span?

是否需要定义一个 trasnaction, 用于串联起整个上下文? 还是 spans 传递? 想清楚如何串联分布式span.

SlowThreshold 操作过慢时自动执行的操作. 所有 hook 都要支持, 可配置 metrics 和 logger.
metrics 名称: http 为路径名, rpc 为方法名, mysql 为表名, redis 为库名.
1. rpc、http: 默认打印日志, 上报 metrics.
2. mysql、redis、kafka: 默认打印日志, 不上报 metrics.
*/
