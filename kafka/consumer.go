package kafka

import (
	"context"
	"time"

	"github.com/gscaffold/helpers/devops"
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	*kafka.Reader
}

func DisconveryConsumer(topic, groupID string, _opts ...OptionConsumer) (*Consumer, error) {
	return DisconveryConsumerAppExclusive("", topic, groupID, _opts...)
}

func DisconveryConsumerAppExclusive(app, topic, groupID string, _opts ...OptionConsumer) (*Consumer, error) {
	addrs, _ := devops.DiscoveryMany(devops.ResourceKafka, app, topic, "")
	cfgs := kafka.ReaderConfig{
		Brokers:        addrs,
		GroupID:        groupID,
		Topic:          topic,
		CommitInterval: time.Second, // sdk 默认同步提交, 这里为了性能改成异步
		Logger:         kafka.LoggerFunc(Infof),
		ErrorLogger:    kafka.LoggerFunc(Errorf),
	}

	for _, opt := range _opts {
		opt(&cfgs)
	}

	if err := cfgs.Validate(); err != nil {
		return &Consumer{}, err
	}

	consumer := kafka.NewReader(cfgs)
	return &Consumer{consumer}, nil
}

// Receive 读取一条消息, 并且自动 commit.
func (c *Consumer) Receive(ctx context.Context) (Message, error) {
	return c.ReadMessage(ctx)
}

func (c *Consumer) ReceiveNotCommit(ctx context.Context) (Message, error) {
	return c.FetchMessage(ctx)
}

func (c *Consumer) Ack(ctx context.Context, msgs ...Message) error {
	return c.CommitMessages(ctx, msgs...)
}
