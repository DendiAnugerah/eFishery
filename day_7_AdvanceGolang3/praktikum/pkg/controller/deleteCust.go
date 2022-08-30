package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h Handler) DeleteCust(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(CustomerList, id)
	return c.NoContent(http.StatusNoContent)
}
