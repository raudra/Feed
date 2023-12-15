package main

import (
	"feed-manager/src/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Recovery())
	router.GET("ping", ping)

	v1 := router.Group("/api/v1")
	{
		v1.POST("feeds", controllers.CreateFeed)
	}

	router.Run()
}

func ping(c *gin.Context) {
	data := map[string]string{
		"status": "ok",
	}
	c.JSON(http.StatusOK, data)
}
