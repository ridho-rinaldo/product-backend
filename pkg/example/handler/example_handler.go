package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ridho-rinaldo/product-backend/libs/models"
	"github.com/ridho-rinaldo/product-backend/pkg/example/model"
)

func (h *HttpHandler) example(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return models.ToJSON(c).BadRequest("Bad Request")
	}

	t := model.ExampleRequest{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		return models.ToJSON(c).BadRequest("Bad Request")
	}

	model := model.ExampleRequest{
		Data: t.Data,
	}

	result, err := h.usecase.Example(model)

	if err != nil {
		resp := &models.Response{Code: 400, Message: err.Error()}
		return c.JSON(http.StatusBadRequest, resp)
	}

	return models.ToJSON(c).Ok(result, "Successfully")
}
