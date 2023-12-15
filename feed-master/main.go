package main

import (
	"feed-master/config"
	"fmt"
)

func main() {
	fmt.Println("Starting Feed Master Service !!!")
	config.InitRedis()
	InitRouter()
}
