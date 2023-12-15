package consumers

import (
	"feed-events/config"
	"log"

	"github.com/IBM/sarama"
)

type ConsumerInterface interface {
	consume()
	getTopicName() string
	setUpConsumer(sarama.Consumer)
	processMessage([]byte)
}

func Start() {
	log.Println("Starting consumers")
	kafkaConfig := config.NewKafkaConfig()
	consumer, _ := sarama.NewConsumer(kafkaConfig.Brokers, nil)
	var feedConsumer ConsumerInterface
	feedConsumer = &FeedConsumer{}
	feedConsumer.setUpConsumer(consumer)
	log.Println("Feed Consumer", feedConsumer)
	go feedConsumer.consume()
	ch := make(chan int64)
	<-ch
}
