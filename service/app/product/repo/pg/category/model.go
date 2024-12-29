package category

import (
	"time"
)

const (
	CATEGORY_TABLE = "komo_category"
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
