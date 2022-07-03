//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/minkj1992/fiber-wire/entity"
)

func InitializeEvent() entity.Event { // injector
	wire.Build(entity.NewEvent, entity.NewGreeter, entity.NewMessage) // provider
	return entity.Event{}
}
