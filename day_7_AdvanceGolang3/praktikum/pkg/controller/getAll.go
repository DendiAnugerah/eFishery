package controller

import (
	"net/http"

	"github.com/DendiA/crud/pkg/model"
	"github.com/labstack/echo/v4"
)

var (
	LastId       = 1
	CustomerList = make(map[int]model.Customers)
)

func (h Handler) GetCust(c echo.Context) error {
	var result []model.Customers
	h.DB.Find(&result)

	// result = append(result, mocks.Custommers...)

	// for key, _ := range mocks.Custommers {
	// 	result = append(result, mocks.Custommers[key])
	// }

	return c.JSON(http.StatusOK, h.DB.Find(&result))
}
