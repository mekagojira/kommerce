package main

import (
	_ "komo/app/auth/controller"
	"komo/lib/db"
	"komo/lib/engine"
	"komo/lib/streaming"
)

func main() {
	db.NewPg(engine.GetEnv("DATABASE_URI", engine.GetEnv("AUTH_DATABASE_URI")))

	streaming.Connect()

	engine.StartServer(engine.GetEnv("AUTH_PORT", "8080"))

}
