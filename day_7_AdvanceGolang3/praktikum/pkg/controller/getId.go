package controller

import (
	"net/http"
	"strconv"

	"github.com/DendiA/crud/pkg/mocks"
	"github.com/labstack/echo/v4"
)

func (h Handler) GetCustbyId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, mocks.Custommers[id])
}
