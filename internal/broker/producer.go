package broker

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
	"time"
)

var brokers = []string{"127.0.0.1:9094"}

type Producer interface {
	Send(topic string, value interface{}) error
	Close()
}

type producer struct {
	prod sarama.SyncProducer
}

func NewProducer() Producer {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	prod, _ := sarama.NewSyncProducer(brokers, config)

	return &producer{
		prod: prod,
	}
}

func (p producer) Send(topic string, value interface{}) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		log.Err(err).Msg("failed marshaling")
	}

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Key:       sarama.StringEncoder(topic),
		Value:     sarama.StringEncoder(bytes),
		Partition: -1,
		Timestamp: time.Time{},
	}

	_, _, err = p.prod.SendMessage(msg)

	return err
}

func (p producer) Close() {
	p.prod.Close()
}
