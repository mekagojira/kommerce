package controller

import (
	"komo/app/auth/feature/ping"
	"komo/lib/engine"
)

func init() {
	engine.RegisterEndpoint("/ping", ping.Handle)
}
