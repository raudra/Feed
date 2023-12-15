package config

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	ddAws "gopkg.in/DataDog/dd-trace-go.v1/contrib/aws/aws-sdk-go/aws"
)

var dbClient *dynamodb.DynamoDB

func InitDb() *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{
		Endpoint:    aws.String("http://host.docker.internal:8000"),
		Region:      aws.String("ap-southeact-1"),
		Credentials: credentials.NewStaticCredentials("test-cred", "test-cred", "test-cred"),
	})

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	ddAws.WrapSession(sess)
	dbClient = dynamodb.New(sess)
	return dbClient
}

func GetDbClient() *dynamodb.DynamoDB {
	if dbClient == nil {
		dbClient = InitDb()
	}
	return dbClient
}
