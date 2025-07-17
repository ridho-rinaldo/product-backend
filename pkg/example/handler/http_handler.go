package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ridho-rinaldo/product-backend/config/postgresql"
	"github.com/ridho-rinaldo/product-backend/pkg/example/usecase"
)

type HttpHandler struct {
	usecase usecase.ExampleUsecase
}

func NewHttpHandler(usecase usecase.ExampleUsecase) *HttpHandler {
	return &HttpHandler{usecase: usecase}
}

func (h *HttpHandler) Mount(g *echo.Group, dbConn *postgresql.DbConnection) {
	g.POST("/example", h.example)
}
