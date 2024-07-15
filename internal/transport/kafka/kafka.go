package kafka

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type Kafka struct {
	KafkaProducer *Producer
	KafkaConsumer *Consumer
}

func NewKafka(url, topic string) *Kafka {
	return &Kafka{
		KafkaProducer: NewProducer(url, topic),
		KafkaConsumer: NewConsumer(url, topic),
	}
}

func EncodeProducerMessage(topic string, value map[string]interface{}) (*sarama.ProducerMessage, error) {
	byteValue, err := json.MarshalIndent(value, "", " ")
	if err != nil {
		logrus.Errorf("Failed to encode producer message: %s", err)
		return nil, err
	}
	return &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(byteValue),
	}, nil
}

func DecodeConsumerMessage(message *sarama.ConsumerMessage) (map[string]interface{}, error) {
	value := make(map[string]interface{})
	if err := json.Unmarshal(message.Value, &value); err != nil {
		logrus.Errorf("Failed to decode consumer message: %s", err)
		return nil, err
	}
	return value, nil
}
