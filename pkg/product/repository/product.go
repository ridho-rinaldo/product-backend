package repository

import "github.com/ridho-rinaldo/product-backend/pkg/product/model"

type ProductRepository interface {
	ListProduct() ([]model.ListProduct, error)
	ProductByID(id string) (model.ListProduct, error)
	NewProduct(payload model.RequestAddProduct) error
	UpdateProduct(payload model.RequestUpdateProduct) error
	DeleteProduct(id string) error
}
