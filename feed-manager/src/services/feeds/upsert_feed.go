package feeds

import (
	"feed-manager/config"
	"feed-manager/src/events/producers"
	"feed-manager/src/models"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func CreateFeed(args map[string]interface{}) (models.Feed, error) {
	feed := models.NewFeed(args)
	feed, err := saveToDb(feed)

	if err != nil {
		log.Fatalf("Save to db failed with error: %s", err)
		return models.Feed{}, err
	} else {
		notifyGenerationEvent(feed)
	}

	return feed, nil
}

func saveToDb(feed models.Feed) (models.Feed, error) {
	client := config.GetDbClient()
	av, err := dynamodbattribute.MarshalMap(feed)

	if err != nil {
		return models.Feed{}, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(config.FeedTable),
	}

	_, err = client.PutItem(input)

	if err != nil {
		return models.Feed{}, err
	}

	log.Println("Feed saved successfully with details %+v", feed)

	return feed, nil
}

func notifyGenerationEvent(feed models.Feed) {
	producer := producers.NewProducer("feed", map[string]interface{}{"sync": false})
	producer.Publish(feed)
}
