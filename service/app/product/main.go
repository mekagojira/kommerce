package main

import (
	"komo/lib/db"
	"komo/lib/engine"
	"komo/lib/streaming"

	_ "komo/app/product/controller"
)

func main() {
	streaming.Connect(streaming.KafkaConfig{ConsumerGroup: engine.GetEnv("KAFKA_CONSUMER_GROUP", "CONSUMER_GROUP")})

	db.NewPg(engine.GetEnv("DATABASE_URI", "PRODUCT_DATABASE_URI"))

	engine.StartServer(engine.GetEnv("PRODUCT_PORT", "PORT", "9000"))
}
