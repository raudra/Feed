package config

//import "github.com/Shopify/sarama"

//var brokers [1]string

type KafkaConfig struct {
	Brokers []string
	//Config  *sarama.Config
}

var kafkaConfig *KafkaConfig

func initBrokers() []string {
	return []string{"kafka:9092"}


func InitKakfaConfig() {
	kafkaConfig = &KafkaConfig{
		Brokers: initBrokers(),
		//		Config:  sarama.NewConfig(),
	}
}

func NewKafkaConfig() *KafkaConfig {
	return kafkaConfig
}
