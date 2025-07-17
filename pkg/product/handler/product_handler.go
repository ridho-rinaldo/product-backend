package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ridho-rinaldo/product-backend/libs/models"
	"github.com/ridho-rinaldo/product-backend/pkg/product/model"
)

func (h *HttpHandler) productList(c echo.Context) error {
	result, err := h.usecase.ListProduct()

	if err != nil {
		resp := &models.Response{Code: 400, Message: err.Error()}
		return c.JSON(http.StatusBadRequest, resp)
	}

	return models.ToJSON(c).Ok(result, "Successfully")
}

func (h *HttpHandler) productById(c echo.Context) error {
	id := c.Param("id")

	result, err := h.usecase.ProductByID(id)

	if err != nil {
		resp := &models.Response{Code: 400, Message: err.Error()}
		return c.JSON(http.StatusBadRequest, resp)
	}

	return models.ToJSON(c).Ok(result, "Successfully")
}

func (h *HttpHandler) addProduct(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return models.ToJSON(c).BadRequest("Bad Request")
	}

	t := model.RequestAddProduct{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		return models.ToJSON(c).BadRequest("Bad Request")
	}

	model := model.RequestAddProduct{
		ProductName: t.ProductName,
		Price:       t.Price,
		Stock:       t.Stock,
	}

	err = h.usecase.NewProduct(model)
	if err != nil {
		resp := &models.Response{Code: 400, Message: err.Error()}
		return c.JSON(http.StatusBadRequest, resp)
	}

	return models.ToJSON(c).Ok(nil, "Successfully")
}

func (h *HttpHandler) updateProduct(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return models.ToJSON(c).BadRequest("Bad Request")
	}

	t := model.RequestUpdateProduct{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		return models.ToJSON(c).BadRequest("Bad Request")
	}

	model := model.RequestUpdateProduct{
		IDProduct:   t.IDProduct,
		ProductName: t.ProductName,
		Price:       t.Price,
		Stock:       t.Stock,
	}

	err = h.usecase.UpdateProduct(model)
	if err != nil {
		resp := &models.Response{Code: 400, Message: err.Error()}
		return c.JSON(http.StatusBadRequest, resp)
	}

	return models.ToJSON(c).Ok(nil, "Successfully")
}

func (h *HttpHandler) deleteProduct(c echo.Context) error {
	id := c.Param("id")

	err := h.usecase.DeleteProduct(id)

	if err != nil {
		resp := &models.Response{Code: 400, Message: err.Error()}
		return c.JSON(http.StatusBadRequest, resp)
	}

	return models.ToJSON(c).Ok(nil, "Successfully")
}
