package models

type StoreChange struct {
	Give_StoreId string `json:"give_store_id"`
	Get_StoreId  string `json:"get_store_id"`
	ProdId       string `json:"prod_id"`
	Count        string `json:"count"`
}

type ReadFrom struct {
	Give_StoreId   int    `json:"store_id"`
	Give_StoreName string `json:" store_name"`
	ProdId         int    `json:"product_id"`
	ProdName       string `json:"product_name"`
	Count          int    `json:"quantity"`
}

type ReadTo struct {
	Get_StoreId   int    `json:"store_id"`
	Get_StoreName string `json:" store_name"`
	ProdId        int    `json:"product_id"`
	ProdName      string `json:"product_name"`
	Count         int    `json:"quantity"`
}

type Valid struct {
	Store int `json:"store_id"`
	Prod  int `json:"product_id"`
}

type GiveMe struct {
	Day string `json:"day"`
}

type Answer struct {
	StaffName string  `json:"employe"`
	Category  string  `json:"category"`
	Product   string  `json:"product"`
	Quantity  int     `json:"quantity"`
	Summ      float32 `json:"summ"`
}

type Promo struct {
	Promo_id    int    `json:"promo_id"`
	PromoName   string `json:"promo_name"`
	IsPercent   bool   `json:"is_percent"`
	Discount    int    `json:"discount"`
	Limit_Price int    `json:"order_limit_price"`
}

type PromoCreate struct {
	PromoName   string `json:"promo_name"`
	IsPercent   bool   `json:"is_percent"`
	Discount    int    `json:"discount"`
	Limit_Price int    `json:"order_limit_price"`
}
type PromoPrimaryKey struct {
	Promo_id int `json:"promo_id"`
}

type Query struct {
	Offset int    `json:"offset" default:"0"`
	Limit  int    `json:"limit" default:"10"`
	Search string `json:"search"`
}

type StigmaApi struct {
	Order_id   string `json:"order_id"`
	Promo_Code string `json:"promo_code"`
}

type SigmaSql struct {
	Order_id   int     `json:"order_id"`
	List_price float64 `json:"list_price"`
	Discount   float64 `json:"discount"`
}

type ReadStockBigQ struct {
	Store_id   int `json:"store_id"`
	Product_id int `json:"product_id"`
	Quantity   int `json:"quantity"`
}
