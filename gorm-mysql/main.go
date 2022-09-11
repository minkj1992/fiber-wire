package main

import (
	"github.com/minkj1992/go-playground/gorm-mysql/persistance"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	persistance.InitSqlite()
}
