//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/minkj1992/go-playground/wire/entity"
)

func InitializeEvent(phrase string) (entity.Event, error) { // injector
	wire.Build(entity.NewEvent, entity.NewGreeter, entity.NewMessage) // provider
	return entity.Event{}, nil
}
