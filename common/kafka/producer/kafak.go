package producer

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Producer is a small interface used by the application code to publish messages.
// Having this interface makes it easier to test the services that depend on it.
type Producer interface {
	Publish(topic string, key string, value []byte) error
	HandleDeliveryReports()
	Close() error
}

type KafkaProducer struct {
	Producer *kafka.Producer
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

// Publish sends a message and waits for the delivery report with a timeout.
// This prevents indefinite blocking during network partitions or broker issues.
func (kp *KafkaProducer) Publish(topic string, key string, value []byte) error {

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          value,
	}

	if err := kp.Producer.Produce(message, nil); err != nil {
		return err
	}

	return nil

}

// Close flushes and closes the underlying producer.
func (kp *KafkaProducer) Close() error {
	if kp == nil || kp.Producer == nil {
		return nil
	}
	// Flush outstanding messages (timeout in ms)
	kp.Producer.Flush(5000)
	kp.Producer.Close()
	return nil
}

func (kp *KafkaProducer) HandleDeliveryReports() {

	for e := range kp.Producer.Events() {

		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Printf("Delivery failed: %v\n", ev.TopicPartition.Error)
			} else {
				fmt.Printf("Message delivered to %v\n", ev.TopicPartition)
			}
		}
	}
}
