package main

import (
	approuters "InsuranceChatWS/internal/app_routers"
	"InsuranceChatWS/internal/configuration"
	"log"
)

func main() {
	container, err := configuration.BuildContainer()
	if err != nil {
		log.Fatalf("Failed to build container: %v", err)
	}

	// Ensure cleanup on shutdown
	defer container.Close()

	// Setup routers
	approuters.StartServer(container)
}
