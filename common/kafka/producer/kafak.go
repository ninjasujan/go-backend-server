package producer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	Producer *kafka.Producer
}

type Producer interface {
	Publish(topic string, key string, value []byte) error
}

func NewKafkaProducer(broker string, clientId string) (*KafkaProducer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"client.id":         clientId,
		"acks":              "all",
	})

	if err != nil {
		return nil, err
	}

	return &KafkaProducer{
		Producer: producer,
	}, nil
}

func (kp *KafkaProducer) Publish(topic string, key string, value []byte) error {

	ch := make(chan kafka.Event, 1)

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          value,
	}

	kp.Producer.Produce(message, ch)

	// Wait for delivery report

	msg := <-ch
	m := msg.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	} else {
		return nil
	}

}
