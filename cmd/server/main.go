package main

import (
	"fmt"
	"github.com/kizuru/passkrypt-server/pkg/unregistering"
	"log"
	"net/http"

	"github.com/kizuru/passkrypt-server/pkg/adding"
	"github.com/kizuru/passkrypt-server/pkg/http/rest"
	"github.com/kizuru/passkrypt-server/pkg/listing"
	"github.com/kizuru/passkrypt-server/pkg/logging"
	"github.com/kizuru/passkrypt-server/pkg/registering"
	"github.com/kizuru/passkrypt-server/pkg/storage/memory"
)

const (
	Memory int = iota
)

func main() {
	storageType := Memory

	var lister listing.Service
	var adder adding.Service
	var logger logging.Service
	var registerer registering.Service
	var unregisterer unregistering.Service

	switch storageType {
	case Memory:
		s := new(memory.Storage)

		lister = listing.NewService(s)
		adder = adding.NewService(s)
		logger = logging.NewService(s)
		registerer = registering.NewService(s)
		unregisterer = unregistering.NewService(s)
	}

	router := rest.Handler(lister, adder, registerer, unregisterer, logger)

	fmt.Println("The passkrypt server is now running: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
