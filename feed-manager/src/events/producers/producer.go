package producers

import (
	appConfig "feed-manager/config"
	"log"
)

type Producer interface {
	Publish(interface{}) (bool, error)
	//	Logger(args map[string]interface{})
	//	Args(args)
}

func NewProducer(pType string, config map[string]interface{}) Producer {
	var producer Producer

	prd, err := appConfig.InitKafkaClient(config["sync"].(bool))

	if err != nil {
		log.Fatal("Error while connecting with kafka %s", err)
		//return Producer{}
	}

	if pType == "feed" {
		producer = NewFeedProducer(prd)

	}

	return producer
}
