package controller

import (
	"komo/app/product/feature/ping"
	"komo/lib/engine"
	"net/http"
)

func init() {
	engine.Server.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		res := ping.Handle()

		if res.IsOk() {
			w.Write([]byte("pong"))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
	})
}
