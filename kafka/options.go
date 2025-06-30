package kafka

import (
	"time"

	"github.com/segmentio/kafka-go"
)

type OptionProducer func(*kafka.Writer)

func OptionProducerAddr(addrs []string) OptionProducer {
	return func(w *kafka.Writer) {
		w.Addr = kafka.TCP(addrs...)
	}
}

func OptionProducerTopic(topic string) OptionProducer {
	return func(w *kafka.Writer) {
		w.Topic = topic
	}
}

func OptionProducerBalancer(balance kafka.Balancer) OptionProducer {
	return func(w *kafka.Writer) {
		w.Balancer = balance
	}
}

func OptionProducerBatchSize(size int) OptionProducer {
	return func(w *kafka.Writer) {
		w.BatchSize = size
	}
}

func OptionProducerReadTimeout(dur time.Duration) OptionProducer {
	return func(w *kafka.Writer) {
		w.ReadTimeout = dur
	}
}

func OptionProducerWriteTimeout(dur time.Duration) OptionProducer {
	return func(w *kafka.Writer) {
		w.WriteTimeout = dur
	}
}

func OptionProducerRequiredAcks(level kafka.RequiredAcks) OptionProducer {
	return func(w *kafka.Writer) {
		w.RequiredAcks = level
	}
}

func OptionProducerAsync() OptionProducer {
	return func(w *kafka.Writer) {
		w.Async = true
	}
}

func OptionProducerWithErrorLog(logger kafka.LoggerFunc) OptionProducer {
	return func(w *kafka.Writer) {
		w.ErrorLogger = logger
	}
}

// --------------------------------- consumer -------------------

type OptionConsumer func(*kafka.ReaderConfig)

func OptionConsumerAddr(addrs []string) OptionConsumer {
	return func(w *kafka.ReaderConfig) {
		w.Brokers = addrs
	}
}

func OptionConsumerGroupID(groupID string) OptionConsumer {
	return func(w *kafka.ReaderConfig) {
		w.GroupID = groupID
	}
}

func OptionConsumerTopic(topic string) OptionConsumer {
	return func(w *kafka.ReaderConfig) {
		w.Topic = topic
	}
}

func OptionConsumerQueueCapacity(cap int) OptionConsumer {
	return func(w *kafka.ReaderConfig) {
		w.QueueCapacity = cap
	}
}

func OptionConsumerSync() OptionConsumer {
	return func(w *kafka.ReaderConfig) {
		w.CommitInterval = 0
	}
}

func OptionConsumerCommitInterval(dur time.Duration) OptionConsumer {
	return func(w *kafka.ReaderConfig) {
		w.CommitInterval = dur
	}
}
