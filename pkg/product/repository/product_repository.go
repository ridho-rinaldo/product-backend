package repository

import (
	"github.com/ridho-rinaldo/product-backend/config/postgresql"
	"github.com/ridho-rinaldo/product-backend/pkg/product/model"
	"github.com/rs/zerolog/log"
)

type productRepository struct {
	dbConn *postgresql.DbConnection
}

func NewProductRepository(dbConn *postgresql.DbConnection) ProductRepository {
	return &productRepository{dbConn: dbConn}
}

func (r *productRepository) ListProduct() ([]model.ListProduct, error) {
	db := r.dbConn.Db
	result := []model.ListProduct{}

	qr := "SELECT p.product_name, p.price, p.stock, p.created_at FROM product p WHERE p.deleted_at IS NULL"
	err := db.Raw(qr).Scan(&result).Error

	if err != nil {
		log.Error().Msg(err.Error())
		return result, err
	}

	return result, err
}

func (r *productRepository) ProductByID(id string) (model.ListProduct, error) {
	db := r.dbConn.Db
	result := model.ListProduct{}

	qr := "SELECT p.product_name, p.price, p.stock, p.created_at FROM product p WHERE p.id_product = ? AND p.deleted_at IS NULL"
	err := db.Raw(qr, id).Scan(&result).Error

	if err != nil {
		log.Error().Msg(err.Error())
		return result, err
	}

	return result, err
}

func (r *productRepository) NewProduct(payload model.RequestAddProduct) error {
	db := r.dbConn.Db

	p := payload
	qr := `INSERT INTO public.product (product_name, price, stock, created_at)
	VALUES (?, ?, ?, now());`
	err := db.Exec(qr, p.ProductName, p.Price, p.Stock).Error

	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	return err
}

func (r *productRepository) UpdateProduct(payload model.RequestUpdateProduct) error {
	db := r.dbConn.Db

	p := payload
	qr := `UPDATE public.product 
	SET product_name = ?,
	price = ?,
	stock = ?,
	updated_at = now()
	WHERE product.id_product = ?`
	err := db.Exec(qr, p.ProductName, p.Price, p.Stock, p.IDProduct).Error

	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	return err
}

func (r *productRepository) DeleteProduct(id string) error {
	db := r.dbConn.Db
	result := model.ListProduct{}

	qr := `UPDATE public.product 
	SET deleted_at = now()
	WHERE product.id_product = ?`
	err := db.Raw(qr, id).Scan(&result).Error

	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	return err
}
