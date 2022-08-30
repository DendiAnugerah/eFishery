package controller

import (
	"log"
	"net/http"

	"github.com/DendiA/crud/pkg/model"
	"github.com/labstack/echo/v4"
)

func (h Handler) CreateCustomer(c echo.Context) error {
	log.Println("masuk ke crate user")
	req := new(model.Customers)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	// user := model.Customer{
	// 	Id:    LastId,
	// 	Name:  req.Name,
	// 	Email: req.Email,
	// }

	// CustomerList[LastId] = user

	user := model.Customers{
		Id:    LastId,
		Name:  req.Name,
		Email: req.Email,
	}
	LastId++
	if result := h.DB.Create(&user); result.Error != nil {
		log.Println("Result Error")
	}

	return c.JSON(http.StatusCreated, user)
}
