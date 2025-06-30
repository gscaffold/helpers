/*
底层基于 kafka-go 实现, 封装资源发现等功能.
默认通过环境变量发现 topic 地址.

producer 默认异步发送, 如需同步请在初始化时设置 Async 为 true.
关闭 producer 进程前, 请务必调用 Close 方法.

为什么选 segmentio/kafka-go. 主要差别是使用起来是否方便, 性能差距不大.
竞品一 Shopify/sarama: 老牌 sdk, 使用方式复杂.
竞品二 confluent-kafka-go: 需要开启 cgo
*/

package kafka
