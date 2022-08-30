package controller

import (
	"net/http"
	"strconv"

	"github.com/DendiA/crud/pkg/model"
	"github.com/labstack/echo/v4"
)

func (h Handler) UpdateCust(c echo.Context) error {
	cust := new(model.Customers)
	var CustAll = []model.Customers{}
	if err := c.Bind(cust); err != nil {
		return err
	}

	Id, _ := strconv.Atoi(c.Param("id"))
	UpdatedCust := model.Customers{
		Id:    CustAll[Id].Id,
		Name:  cust.Name,
		Email: cust.Email,
	}

	CustAll[Id] = UpdatedCust

	return c.JSON(http.StatusOK, CustAll[Id])
}
