package main

import (
	"komo/lib/db"
	"komo/lib/engine"
	"komo/lib/streaming"

	"komo/app/product/common/constant"
	_ "komo/app/product/controller"
)

func main() {
	streaming.Connect(constant.GetKafkaEnv())

	db.NewPg(constant.GetPgEnv())

	engine.StartServer(constant.GetPortEnv())
}
