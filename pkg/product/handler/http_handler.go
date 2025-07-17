package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ridho-rinaldo/product-backend/config/postgresql"
	"github.com/ridho-rinaldo/product-backend/pkg/product/usecase"
)

type HttpHandler struct {
	usecase usecase.ProductUsecase
}

func NewHttpHandler(usecase usecase.ProductUsecase) *HttpHandler {
	return &HttpHandler{usecase: usecase}
}

func (h *HttpHandler) Mount(g *echo.Group, dbConn *postgresql.DbConnection) {
	g.GET("/product/list", h.productList)
	g.GET("/product/:id", h.productById)
	g.POST("/product", h.addProduct)
	g.PUT("/product", h.updateProduct)
	g.PUT("/product/:id", h.deleteProduct)
}
