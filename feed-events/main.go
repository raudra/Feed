package main

import (
	"feed-events/config"
	"feed-events/src/events"
	"fmt"
)

func main() {
	fmt.Println("Starting Feed Events Service !!!")
	config.InitKakfaConfig()
	events.Start()
}
