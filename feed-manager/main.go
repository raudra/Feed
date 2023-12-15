package main

import (
	"feed-manager/config"
	"fmt"
)

func main() {
	fmt.Println("Starting Feed Manager Service !!!")
	InitRouter()
	config.InitDb()
}
