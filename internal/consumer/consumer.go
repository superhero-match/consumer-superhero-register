package consumer

import (
	"github.com/consumer-superhero-register/internal/config"
	"time"

	"github.com/segmentio/kafka-go"
)

// Consumer holds Kafka consumer related data.
type Consumer struct {
	Consumer *kafka.Reader
}

// NewConsumer configures Kafka consumer that consumes from configured topic.
func NewConsumer(cfg *config.Config) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        cfg.Consumer.Brokers,
		GroupID:        cfg.Consumer.GroupID,
		Topic:          cfg.Consumer.Topic,
		QueueCapacity:  int(1),
		MaxWait:        time.Second,
	})

	return &Consumer{
		Consumer: r,
	}
}