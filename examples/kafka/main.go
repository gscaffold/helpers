package main

import (
	"context"
	"fmt"

	"github.com/gscaffold/helpers/devops"
	"github.com/gscaffold/helpers/kafka"
	"github.com/gscaffold/helpers/logger"
	"github.com/gscaffold/utils"
)

func main() {
	err := devops.Register(devops.ResourceKafka, "", topic, "", []string{"127.0.0.1:9092"})
	utils.HandleFatalError(err, "init", "register kafka dsn error")

	utils.SafeGo(producerFunc, func(_ error) {})
	consumerFunc()
}

var (
	topic = "data.test.example"
)

func producerFunc() {
	// 异步发送增加效率
	producer, err := kafka.DisconveryProducer(topic, kafka.OptionProducerAsync())
	utils.HandleFatalError(err, "producer", "producer init error")
	for i := 0; i < 100; i++ {
		producer.Send(context.Background(), kafka.Message{
			Value: []byte(fmt.Sprintf("no.%d", i)),
		})
	}
	producer.Close(context.TODO())
	logger.Info(context.TODO(), "producer finish")
}

func consumerFunc() {
	consumer, err := kafka.DisconveryConsumer(topic, "example")
	utils.HandleFatalError(err, "consumer", "consumer init error")
	for {
		msg, err := consumer.Receive(context.TODO())
		if err != nil {
			logger.Error(context.TODO(), "consumer error", err)
			continue
		}
		logger.Infof(context.TODO(), "consumer msg:%s", string(msg.Value))
	}
}
