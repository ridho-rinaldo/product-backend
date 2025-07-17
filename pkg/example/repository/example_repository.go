package repository

import "github.com/ridho-rinaldo/product-backend/pkg/example/model"

type ExampleRepository interface {
	Example(payload model.ExampleRequest) (model.ExampleResponse, error)
}
