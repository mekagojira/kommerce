package ping

import (
	"fmt"
	"komo/lib/db"
	"komo/lib/engine"
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

		if res.IsOk() {
			w.Write([]byte("pong"))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
	})
}
