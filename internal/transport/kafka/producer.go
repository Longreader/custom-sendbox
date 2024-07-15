package kafka

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
	"strings"
)

type Producer struct {
	URLs    string
	Topic   string
	Packets chan map[string]interface{}

	ctx    context.Context
	cancel context.CancelFunc
}

func NewProducer(brokers, topic string) *Producer {
	return &Producer{
		URLs:  brokers,
		Topic: topic,
	}
}

func (p *Producer) Run() (chan map[string]interface{}, error) {
	p.Packets = make(chan map[string]interface{})
	p.ctx, p.cancel = context.WithCancel(context.Background())

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewAsyncProducer(strings.Split(p.URLs, ","), config)
	if err != nil {
		logrus.Errorf("Failed to Run Sarama producer: %v", err)
	}

	go p.run(producer)

	return p.Packets, nil
}

func (p *Producer) run(client sarama.AsyncProducer) {
ProducerLoop:
	for {
		select {
		case message := <-p.Packets:
			saramaMessage, err := EncodeProducerMessage(p.Topic, message)
			if err != nil {
				logrus.Errorf("Failed to decode message: %v", err)
				continue
			}
			client.Input() <- saramaMessage
		case err := <-client.Errors():
			if err != nil {
				logrus.Errorf("Failed to produce message: %v", err)
			}
		case success := <-client.Successes():
			if success != nil {
				logrus.Debugf("Produced message to topic %s, with value %s \n", success.Topic, success.Value)
			}
		case <-p.ctx.Done():
			break ProducerLoop
		}
	}
	logrus.Infof("Closing Sarama producer on topic: %s", p.Topic)
	return
}

func (p *Producer) Stop() {
	p.cancel()

	p.ctx = nil
	p.cancel = nil
}
