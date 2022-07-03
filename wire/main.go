package main

import (
	"fmt"
	"os"

	"github.com/minkj1992/go-playground/wire/entity"
)

func beforeWire() {
	const msg = "This is code without wire"
	message := entity.NewMessage(msg)
	greeter := entity.NewGreeter(message)
	event, _ := entity.NewEvent(greeter)

	event.Start()
}

func afterWire() {
	const msg = "Hello Dependency injection with wire"
	e, err := InitializeEvent(msg)
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()
}

func main() {
	// beforeWire()
	afterWire()
}
