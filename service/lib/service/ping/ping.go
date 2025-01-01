package ping

import (
	"context"
	"fmt"
	"komo/lib/db"
	"komo/lib/engine"
	"komo/lib/streaming"
	"net/http"
)

func Ping() {
	engine.Server.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		res := engine.NewResult(true)

		row := db.Db.QueryRowBg(`SELECT 1 + 1 as result`)

		if !row.IsOk() {
			engine.Logger.Error(row.Error.Error())
			res.WithError(row.Error)
		}
		dbResult := 0
		row.PureData().Scan(&dbResult)
		if dbResult != 2 {
			res.WithError(fmt.Errorf("DB error"))
		}

		if !res.IsOk() {
			sendError(w, r)
			engine.Logger.Error(res.Error.Error())
			return
		}

		err := streaming.Kafka.Ping(context.Background())
		if err != nil {
			engine.Logger.Error(err.Error())
			res.WithError(err)
			sendError(w, r)
			return
		}

		sendOk(w, r)
	})
}

func sendOk(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func sendError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}
