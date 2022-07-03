package main

import "github.com/minkj1992/fiber-wire/entity"

func beforeWire() {
	message := entity.NewMessage()
	greeter := entity.NewGreeter(message)
	event := entity.NewEvent(greeter)

	event.Start()
}

func afterWire() {
	e := InitializeEvent()
	e.Start()
}

func main() {
	// beforeWire()
	afterWire()
}
