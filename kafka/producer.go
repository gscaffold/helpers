package kafka

import (
	"context"
	"errors"

	"github.com/gscaffold/helpers/devops"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	*kafka.Writer
}

func DisconveryProducer(topic string, _opts ...OptionProducer) (*Producer, error) {
	return DisconveryProducerAppExclusive("", topic, _opts...)
}

func DisconveryProducerAppExclusive(app, topic string, _opts ...OptionProducer) (*Producer, error) {
	addrs, _ := devops.DiscoveryMany(devops.ResourceKafka, app, topic, "")
	producer := &kafka.Writer{
		Addr:        kafka.TCP(addrs...),
		Topic:       topic,
		Balancer:    &kafka.LeastBytes{},
		Logger:      kafka.LoggerFunc(Infof),
		ErrorLogger: kafka.LoggerFunc(Errorf),
		// Transport
	}

	for _, opt := range _opts {
		opt(producer)
	}

	if len(producer.Addr.String()) == 0 {
		return &Producer{}, errors.New("no kafka address")
	}

	return &Producer{producer}, nil
}

func (p *Producer) Send(ctx context.Context, msgs ...Message) error {
	err := p.Writer.WriteMessages(ctx, msgs...)
	return err
}

func (p *Producer) Close(ctx context.Context) error {
	return p.Writer.Close()
}
