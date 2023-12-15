package config

import (
	"github.com/Shopify/sarama"
)

//var brokers [1]string

func InitKafkaClient(sync bool) (interface{}, error) {
	brokers := initBrokers()

	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	var prd interface{}
	var err error

	if sync {
		prd, err = sarama.NewSyncProducer(brokers, config)
	} else {
		prd, err = sarama.NewAsyncProducer(brokers, config)
	}

	return prd, err
}

func initBrokers() []string {
	return []string{"kafka:9092"}
}
