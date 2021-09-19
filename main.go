package main

import (
	"go-assesment/migration"
	"go-assesment/route"
)

func main() {
	migration.Migrate()
	route.Init()
}
