package repository

import (
	"github.com/ridho-rinaldo/product-backend/config/postgresql"
	"github.com/ridho-rinaldo/product-backend/pkg/example/model"
	"github.com/rs/zerolog/log"
)

type exampleRepository struct {
	dbConn *postgresql.DbConnection
}

func NewExampleRepository(dbConn *postgresql.DbConnection) ExampleRepository {
	return &exampleRepository{dbConn: dbConn}
}

func (r *exampleRepository) Example(payload model.ExampleRequest) (model.ExampleResponse, error) {
	db := r.dbConn.Db
	result := model.ExampleResponse{}

	qr := "SELECT * FROM f_example(?)"
	err := db.Raw(qr, payload.Data).Scan(&result).Error

	if err != nil {
		log.Error().Msg(err.Error())
		return result, err
	}

	return result, err
}
