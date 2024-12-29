package product

import "time"

const (
	PRODUCT_TABLE = "komo_product"
)

type ProductVariantData struct {
	Id          string                 `json:"id"`
	ProductId   string                 `json:"productId"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Price       string                 `json:"price"`
	Options     []ProductVariantOption `json:"options"`
	Images      []ImageData            `json:"images"`
	CreatedAt   time.Time              `json:"createdAt"`
	UpdatedAt   time.Time              `json:"updatedAt"`
}

type ImageData struct {
	Path string `json:"path"`
	Name string `json:"name"`
	Mime string `json:"mime"`
}

type ProductVariantOption struct {
	ProductOptionId string `json:"productOptionId"`
	Name            string `json:"name"`
	State           string `json:"state"`
}

type ProductOption struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Variant bool   `json:"variant"`
	State   string `json:"state"`
}

type ProductData struct {
	Id          string                 `json:"id"`
	Slug        string                 `json:"slug"`
	State       string                 `json:"state"`
	Name        string                 `json:"name"`
	Price       string                 `json:"price"`
	Description string                 `json:"description"`
	Images      []ImageData            `json:"images"`
	Options     []ProductVariantOption `json:"options"`
	Categories  []string               `json:"categories"`
	Tags        []string               `json:"tags"`
	CreatedAt   time.Time              `json:"createdAt"`
	UpdatedAt   time.Time              `json:"updatedAt"`
}

type ProductRow struct {
	Id        string    `db:"id"`
	Slug      string    `db:"slug"`
	Name      string    `db:"name"`
	State     string    `db:"state"`
	Price     string    `db:"price"`
	Data      []byte    `db:"data"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
