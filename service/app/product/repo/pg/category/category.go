package category

import (
	"fmt"
	constant "komo/app/product/common/constant"
	"komo/app/product/repo/pg"
	"komo/lib/db"
	"komo/lib/engine"
)

func CategorySlugExists(slug string) *engine.Result[bool] {
	res := engine.NewResult(true)

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
	res := engine.NewResult(true)

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

type ListCategoriesInput struct {
	State string
}

func ListCategories(input ListCategoriesInput, paging pg.Paging) *engine.Result[[]CategoryData] {
	res := engine.NewResult([]CategoryData{})

	query := map[string]string{}
	params := map[string]any{}

	if input.State != "" {
		query["state"] = "state = $1"
		params["state"] = input.State
	} else {
		query["state"] = "state != $1"
		params["state"] = constant.CATEGORY_STATE_DELETED
	}

	if paging.LastId != "" {
		query["last_id"] = "slug > $2"
		params["last_id"] = paging.LastId
	} else {
		query["last_id"] = "1 = $2"
		params["last_id"] = "1"
	}

	query["limit"] = "limit $3"
	params["limit"] = paging.Limit

	sql := `SELECT slug, category_name, state, created_at FROM ` + CATEGORY_TABLE + ` WHERE ` + query["state"] + ` AND ` + query["last_id"] + ` ORDER BY slug ASC ` + query["limit"]

	rows := db.Db.QueryBg(sql, params["state"], params["last_id"], params["limit"])

	if !rows.IsOk() {
		return res.WithError(rows.Error)
	}
	if rows.Data == nil {
		return res.WithError(fmt.Errorf("no data"))
	}

	var result []CategoryData
	for rows.PureData().Next() {
		var row CategoryRow
		rows.PureData().Scan(&row.Slug, &row.CategoryName, &row.State, &row.CreatedAt)
		result = append(result, CategoryData{
			Slug:         row.Slug,
			CategoryName: row.CategoryName,
			State:        row.State,
			CreatedAt:    row.CreatedAt,
		})
	}

	res.WithData(&result)

	return res
}
