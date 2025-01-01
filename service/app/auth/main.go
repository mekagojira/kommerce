package main

import (
	_ "komo/app/auth/controller"
	"komo/lib/db"
	"komo/lib/engine"
	"komo/lib/streaming"
)

func main() {
	streaming.Connect(streaming.KafkaConfig{ConsumerGroup: engine.GetEnv("AUTH_CONSUMER_GROUP", "CONSUMER_GROUP")})

	db.NewPg(engine.GetEnv("DATABASE_URI", "AUTH_DATABASE_URI"))

	engine.StartServer(engine.GetEnv("AUTH_PORT", "PORT", "9000"))
}
