package consumers

import (
	"encoding/json"
	"feed-events/src/models"
	"feed-events/src/services"

	"log"

	"github.com/IBM/sarama"
)

const (
	FeedTopicName = "feed-producer-topic"
)

type FeedConsumer struct {
	consumer  sarama.Consumer
	topicName string
	groupId   string
}

func (self *FeedConsumer) consume() {
	topic := self.getTopicName()
	log.Println("Starting Consuming messages", self)

	partitionList, _ := self.consumer.Partitions(topic) //get all partitions
	initialOffset := sarama.OffsetOldest                //offset to start reading message from
	for _, partition := range partitionList {
		pc, _ := self.consumer.ConsumePartition(topic, partition, initialOffset)
		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				self.processMessage(message.Value)
			}
		}(pc)
	}
}

func (self FeedConsumer) getTopicName() string {
	return FeedTopicName
}

func (self *FeedConsumer) setUpConsumer(kClient sarama.Consumer) {
	self.consumer = kClient
	self.topicName = self.getTopicName()
}

func (sefl *FeedConsumer) processMessage(msg []byte) {
	feed := &models.Feed{}
	_ = json.Unmarshal(msg, feed)
	services.UpsertFeedNotify(feed)
}
