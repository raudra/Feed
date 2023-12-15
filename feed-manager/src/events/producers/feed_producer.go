package producers

import (
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
)

const (
	FeedProducerTopic = "feed-producer-topic"
)

type Feed struct {
	Topic    string
	Msg      map[string]interface{}
	Producer interface{}
}

func NewFeedProducer(producer interface{}) Feed {
	return Feed{
		Topic:    FeedProducerTopic,
		Producer: producer,
	}
}

func (self Feed) Publish(msg interface{}) (bool, error) {

	producer := self.Producer.(sarama.AsyncProducer)

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	data, _ := json.Marshal(msg)

	message := &sarama.ProducerMessage{
		Topic:     FeedProducerTopic,
		Partition: -1,
		Value:     sarama.StringEncoder(data),
	}

	producer.Input() <- message
	return true, nil

}
