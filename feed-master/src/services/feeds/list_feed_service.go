package feeds

import (
	"context"
	"encoding/json"
	"feed-master/config"
	"feed-master/src/models"
	"fmt"
)

func ListFeed() []models.Feed {
	feeds := fetchFromCache(1)
	return feeds

}

func fetchFromCache(userId int64) []models.Feed {
	ctx := context.TODO()

	redisClient := config.RedisClient()

	list, _ := redisClient.LRange(ctx, fmt.Sprintf("users:%d", userId), 0, -1).
		Result()

	feeds := []models.Feed{}

	for _, ele := range list {
		data, err := redisClient.Get(ctx, ele).
			Result()

		if err == nil {
			feed := models.Feed{}
			_ = json.Unmarshal([]byte(data), &feed)
			feeds = append(feeds, feed)
		}

	}
	return feeds
}
