package events

import (
	"feed-events/src/events/consumers"
)

func Start() {
	//	config.Init
	consumers.Start()
}
