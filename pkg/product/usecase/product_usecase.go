package usecase

import (
	"github.com/ridho-rinaldo/product-backend/pkg/product/model"
	"github.com/ridho-rinaldo/product-backend/pkg/product/repository"
)

type productUsecase struct {
	repo repository.ProductRepository
}

func NewProductUsecase(repository repository.ProductRepository) ProductUsecase {
	return &productUsecase{repo: repository}
}

func (u *productUsecase) ListProduct() ([]model.ListProduct, error) {
	result, err := u.repo.ListProduct()

	return result, err
}

func (u *productUsecase) ProductByID(id string) (model.ListProduct, error) {
	result, err := u.repo.ProductByID(id)

	return result, err
}

func (u *productUsecase) NewProduct(payload model.RequestAddProduct) error {
	err := u.repo.NewProduct(payload)

	return err
}

func (u *productUsecase) UpdateProduct(payload model.RequestUpdateProduct) error {
	err := u.repo.UpdateProduct(payload)

	return err
}

func (u *productUsecase) DeleteProduct(id string) error {
	err := u.repo.DeleteProduct(id)

	return err
}
