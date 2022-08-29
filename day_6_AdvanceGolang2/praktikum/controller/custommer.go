package controller

import (
	"log"
	"net/http"
	"strconv"

	"echo-crud/model"

	"github.com/labstack/echo/v4"
)

var (
	LastId       = 1
	CustomerList = make(map[int]model.Customer)
)

func GetCust(c echo.Context) error {
	var result []model.Customer

	result = append(result, model.Customer{})

	for key, _ := range CustomerList {
		result = append(result, CustomerList[key])
	}

	log.Println("Showing all customer")
	return c.JSON(http.StatusOK, result)
}

func CreateCustomer(c echo.Context) error {
	req := new(model.Customer)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	user := model.Customer{
		Id:    LastId,
		Name:  req.Name,
		Email: req.Email,
	}

	CustomerList[LastId] = user
	LastId++

	log.Println("Customer created")
	return c.JSON(http.StatusCreated, user)
}

func GetCustbyId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	log.Println("Showing cust: ", id)
	return c.JSON(http.StatusOK, CustomerList[id])
}

func UpdateCust(c echo.Context) error {
	cust := new(model.Customer)
	var CustAll = []model.Customer{}
	if err := c.Bind(cust); err != nil {
		return err
	}

	Id, _ := strconv.Atoi(c.Param("id"))
	UpdatedCust := model.Customer{
		Id:    CustAll[Id].Id,
		Name:  cust.Name,
		Email: cust.Email,
	}

	CustAll[Id] = UpdatedCust

	return c.JSON(http.StatusOK, CustAll[Id])
}

func DeleteCust(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(CustomerList, id)
	log.Println("Customer deleted")
	return c.NoContent(http.StatusNoContent)
}
