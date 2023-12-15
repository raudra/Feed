package controllers

import (
	FeedService "feed-master/src/services/feeds"
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Success bool        `json:"success"`
	Err     interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data, omitempty"`
}

func FeedList(c *gin.Context) {
	resp := response{
		Success: true,
		Data:    FeedService.ListFeed(),
	}
	c.JSON(http.StatusOK, resp)
}
