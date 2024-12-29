package product

import (
	"komo/lib/db"
	"komo/lib/engine"
	"time"
)

func CreateProduct(data ProductData) *engine.Result[bool] {
	res := engine.NewResult(true)
	now := time.Now()

	if err := engine.SetUid(&data.Id); !err.IsOk() {
		db.Logger.Error(err.Error.Error())
		return res.WithError(err.Error)
	}

	data.CreatedAt = now
	data.UpdatedAt = now

	jsonData := engine.StructToJsonBytes(data)
	if jsonData.Error != nil {
		return res.WithError(jsonData.Error)
	}

	row := ProductRow{
		Slug:      data.Slug,
		Name:      data.Name,
		State:     data.State,
		Price:     data.Price,
		Data:      *jsonData.Data,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	sql := `
		INSERT INTO ` + PRODUCT_TABLE + ` (id, slug, name, state, price, data, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	db.Db.ExecBg(sql, row.Id, row.Slug, row.Name, row.State, row.Price, row.Data, row.CreatedAt, row.UpdatedAt)

	return res
}
