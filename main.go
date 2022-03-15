package main

import (
	"github.com/Eydzhpee08/university/handlers/db"
	"github.com/Eydzhpee08/university/handlers/routers"
)

func main() {

	// data migrate function call
	db.DataMigration()
	// handler function call
	routers.HandlerRouting()
}
