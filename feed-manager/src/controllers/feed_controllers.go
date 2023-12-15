package controllers

import (
	feedService "feed-manager/src/services/feeds"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status bool        `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Err    interface{} `json:"error,omitempty"`
}

func CreateFeed(c *gin.Context) {
	args := make(map[string]interface{})

	if err := c.BindJSON(&args); err != nil {
		log.Fatalln(err)
	}

	feed, _ := feedService.CreateFeed(args)

	resp := Response{
		Status: true,
		Data: map[string]interface{}{
			"feed": feed,
		},
	}

	c.JSON(http.StatusOK, resp)
}
