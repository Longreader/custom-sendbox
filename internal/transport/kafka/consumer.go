package kafka

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
	"strings"
	"sync"
)

type Packet map[string]interface{}

type Consumer struct {
	URLs    string
	Topic   string
	Packets chan map[string]interface{}

	mutex  sync.Mutex
	ctx    context.Context
	cancel context.CancelFunc
}

func NewConsumer(brokers, topic string) *Consumer {
	return &Consumer{
		URLs:    brokers,
		Topic:   topic,
		Packets: make(chan map[string]interface{}),
	}
}

func (c *Consumer) Run() (chan map[string]interface{}, error) {
	c.Packets = make(chan map[string]interface{})
	c.ctx, c.cancel = context.WithCancel(context.Background())

	config := sarama.NewConfig()
	config.Version = sarama.V3_6_0_0

	client, err := sarama.NewConsumer(strings.Split(c.URLs, ","), config)
	if err != nil {
		logrus.Errorf("Failed to Run Sarama consumer: %v", err)
		return nil, err
	}

	go c.run(client)

	return c.Packets, nil
}

func (c *Consumer) run(client sarama.Consumer) {
	defer func() {
		if err := client.Close(); err != nil {
			logrus.Errorf("Error closing partition consumer: %v", err)
		}
	}()

	partitionConsumer, err := client.ConsumePartition(c.Topic, 0, sarama.OffsetNewest)
	if err != nil {
		logrus.Errorf("Error create consumer: %v", err)
		return
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			logrus.Errorf("Error close partition consumer: %v", err)
			return
		}
	}()

ConsumerLoop:
	for {
		select {
		case message, ok := <-partitionConsumer.Messages():
			c.mutex.Lock()
			logrus.Debugf("received message")
			if !ok {
				logrus.Errorf("message channel was closed")
				return
			}

			packet, err := DecodeConsumerMessage(message)
			if err != nil {
				logrus.Errorf("failed to decode message: %v", err)
				continue
			}

			//"device_id": "string",
			//"device_name": "string",
			//"file_id": "string",
			//"protocol": "string",
			//"vendor": "string"
			//if packet["device_id"] == nil || packet["device_name"] == nil {
			//	logrus.Errorf("packet is missing required fields")
			//	continue
			//}
			//if packet["vendor"] == nil || packet["file_id"] == nil || packet["protocol"] == nil {
			//	logrus.Errorf("packet is missing required fields")
			//	continue
			//}

			if packet["message"] == nil {
				logrus.Errorf("packet is missing required fields")
				continue
			}

			c.Packets <- packet
			c.mutex.Unlock()
		case <-c.ctx.Done():
			break ConsumerLoop
		}
	}
	return
}

func (c *Consumer) Stop() {
	c.cancel()

	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.ctx = nil
	c.cancel = nil
}
