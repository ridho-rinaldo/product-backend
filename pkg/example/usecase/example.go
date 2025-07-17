package usecase

import (
	"github.com/ridho-rinaldo/product-backend/pkg/example/model"
	"github.com/ridho-rinaldo/product-backend/pkg/example/repository"
)

type authExample struct {
	repo repository.ExampleRepository
}

func NewExampleUsecase(repo repository.ExampleRepository) ExampleUsecase {
	return &authExample{repo: repo}
}

func (u *authExample) Example(payload model.ExampleRequest) (model.ExampleResponse, error) {
	result, err := u.repo.Example(payload)

	return result, err
}
