package main

import (
	_ "komo/app/auth/controller"
	"komo/lib/db"
	"komo/lib/engine"
)

func main() {
	db.NewPg(engine.GetEnv("DATABASE_URI", engine.GetEnv("AUTH_DATABASE_URI")))

	engine.StartServer(engine.GetEnv("AUTH_PORT", "8080"))
}
