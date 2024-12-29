package ping

import (
	"fmt"
	"komo/lib/db"
	"komo/lib/engine"
)

func Handle() *engine.Result[bool] {
	res := engine.NewResult(true)

	row := db.Db.QueryRowBg(`SELECT 1 + 1 as result`)
	if !row.IsOk() {
		engine.Logger.Error(row.Error.Error())
		return res.WithError(row.Error)
	}
	dbResult := 0
	row.PureData().Scan(&dbResult)
	if dbResult != 2 {
		return res.WithError(fmt.Errorf("DB error"))
	}

	return res
}
