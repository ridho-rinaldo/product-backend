package model

type ListProduct struct {
	ProductName string `json:"product_name"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
	CreatedAt   string `json:"created_at"`
}

type RequestAddProduct struct {
	ProductName string `json:"product_name"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
}

type RequestUpdateProduct struct {
	IDProduct   string `json:"id_product"`
	ProductName string `json:"product_name"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
}
