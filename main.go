package main

import (
	"SandBox/internal/transport/kafka"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	TEST_TOPIC = "test-topic"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	urls := "10.0.0.23:29094, 10.0.0.23:29095"

	OneKafka := kafka.NewKafka(urls, TEST_TOPIC)
	SecondKafka := kafka.NewKafka(urls, TEST_TOPIC)

	firstChatIn, err := OneKafka.KafkaProducer.Run()
	if err != nil {
		logrus.Fatal(err)
	}
	secondChatIn, err := SecondKafka.KafkaProducer.Run()
	if err != nil {
		logrus.Fatal(err)
	}

	firstChatOut, err := OneKafka.KafkaConsumer.Run()
	if err != nil {
		logrus.Fatal(err)
	}
	secondChatOut, err := SecondKafka.KafkaConsumer.Run()
	if err != nil {
		logrus.Fatal(err)
	}

	go func() {
		for {
			logrus.Info("Start")
			select {
			case msg := <-firstChatOut:
				logrus.Infof("First chat: %s", msg)
			case msg := <-secondChatOut:
				logrus.Infof("Second chat: %s", msg)
			}
		}
	}()

	go func() {
		for {
			time.Sleep(5 * time.Second)
			firstChatIn <- map[string]interface{}{
				"message": "Hello from first chat",
			}
		}
	}()

	go func() {
		for {
			time.Sleep(5 * time.Second)
			secondChatIn <- map[string]interface{}{
				"message": "Hello from second chat",
			}
		}
	}()

	time.Sleep(20 * time.Second)
}
