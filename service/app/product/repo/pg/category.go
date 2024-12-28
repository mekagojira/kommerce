package pg

import (
	"fmt"
	"komo/lib/db"
	"komo/lib/engine"
	"time"
)

const (
	CATEGORY_TABLE = "komo_category"
)

const (
	CATEGORY_STATE_ACTIVE   = "ACTIVE"
	CATEGORY_STATE_INACTIVE = "INACTIVE"
	CATEGORY_STATE_DELETED  = "DELETED"
)

type CategoryRow struct {
	Slug         string    `db:"slug"`
	CategoryName string    `db:"category_name"`
	State        string    `db:"state"`
	CreatedAt    time.Time `db:"created_at"`
	Data         []byte    `db:"data"`
}

type CategoryData struct {
	Slug         string    `json:"slug"`
	CategoryName string    `json:"categoryName"`
	State        string    `json:"state"`
	CreatedAt    time.Time `json:"createdAt"`
}

func CategorySlugExists(slug string) *engine.Result[bool] {
	res := engine.NewResult[bool]()

	rows := db.Db.QueryRowBg(`
		SELECT count(slug) as s FROM `+CATEGORY_TABLE+` WHERE slug = $1`, slug)

	if !rows.IsOk() {
		return res.WithError(res.Error)
	}
	if rows.Data == nil {
		return res.WithError(fmt.Errorf("no data"))
	}

	var count int
	rows.PureData().Scan(&count)

	return res.WithPureData(count > 0)
}

func CreateCategory(data CategoryData) *engine.Result[bool] {
	res := engine.NewResult[bool]()

	jsonData := engine.StructToJsonBytes(data)
	if jsonData.Error != nil {
		return res.WithError(jsonData.Error)
	}

	row := CategoryRow{
		Slug:         data.Slug,
		CategoryName: data.CategoryName,
		State:        data.State,
		CreatedAt:    data.CreatedAt,
		Data:         *jsonData.Data,
	}

	db.Db.ExecBg(`
		INSERT INTO `+CATEGORY_TABLE+` (slug, category_name, state, created_at, data)
		VALUES ($1, $2, $3, $4, $5)
	`, row.Slug, row.CategoryName, row.State, row.CreatedAt, row.Data)

	return res
}
