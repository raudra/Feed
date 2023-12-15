package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	FeedTable = "feeds"
)

type Feed struct {
	UUID      string `json:"uuid, omitempty"`
	UserId    int64  `json:"user_id, omitempty"`
	Title     string `json:"title, omitempty"`
	Desc      string `json:"description, omitempty"`
	CreatedAt int64  `json:"created_at, omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

func NewFeed(args map[string]interface{}) Feed {
	return Feed{
		UUID:      (uuid.New()).String(),
		UserId:    int64(args["user_id"].(float64)),
		Title:     args["title"].(string),
		Desc:      args["description"].(string),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
}
