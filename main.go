package main

import (
	"github.com/thisiserico/golabox/apiservice"
	"github.com/thisiserico/golabox/domain/service"
	"github.com/thisiserico/golabox/eventbus"
	"github.com/thisiserico/golabox/eventbus/publisher"
	"github.com/thisiserico/golabox/eventbus/subscriber"
	"github.com/thisiserico/golabox/memoryds"
)

func main() {
	// Set up databases and event bus
	wc := memoryds.NewWriteClient()
	rc := memoryds.NewReadClient()
	bus := make(chan eventbus.Event, 10)

	// Set up domain
	publisher := publisher.NewPublisher(bus)
	domainService := service.New(wc, wc, rc, publisher)
	subscriber := subscriber.NewSubscriber(bus, domainService)
	go subscriber.HandleEvents()

	// Run server
	apiClient := apiservice.New(domainService, rc)
	apiClient.Run()
}
