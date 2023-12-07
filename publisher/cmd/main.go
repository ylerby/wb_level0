package main

import (
	"publisher/pub"
)

func main() {
	publisher := pub.New("test-cluster", "client-2")
	publisher.PublicMessage()
}
