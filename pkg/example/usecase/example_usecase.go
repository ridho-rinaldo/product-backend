package usecase

import "github.com/ridho-rinaldo/product-backend/pkg/example/model"

type ExampleUsecase interface {
	Example(model.ExampleRequest) (model.ExampleResponse, error)
}
