package services

import (
	"context"
	"encoding/json"
	"feed-events/config"
	"feed-events/src/models"
	"fmt"
	"log"
)

func UpsertFeedNotify(feed *models.Feed) {
	upsertRedisClient(feed)
}

func upsertRedisClient(feed *models.Feed) {
	log.Printf("Upserting record feed upsert record in redis %#v", feed)
	redisClient := config.RedisClient()

	msg, _ := json.Marshal(feed)

	ctx := context.TODO()

	key := fmt.Sprintf("feedId:%d", feed.UserId)

	err := redisClient.LPush(ctx, key, msg, 0).Err()
	if err != nil {
		log.Fatal("Error while setting key to redis %s", err)
	}
}
