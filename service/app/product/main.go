package main

import (
	"komo/lib/db"
	"komo/lib/engine"

	_ "komo/app/product/controller"
)

func main() {
	db.NewPg(engine.GetEnv("DATABASE_URI", engine.GetEnv("PRODUCT_DATABASE_URI")))

	engine.StartServer(engine.GetEnv("PRODUCT_PORT", "8080"))
}
