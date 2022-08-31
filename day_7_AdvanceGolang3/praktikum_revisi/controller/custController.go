package controller

import (
	"custommer-crud/infrastructure/config"
	"custommer-crud/model"
	"custommer-crud/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CustomerController struct {
	interact usecase.CustInteract
}

func ConnUser(dbhandler config.DbHandler) *CustomerController {
	return &CustomerController{
		interact: usecase.CustInteract{
			CustRepo: &usecase.CustRepository{
				DbHandler: dbhandler,
			},
		},
	}
}

// POST | create Customer
func (controller *CustomerController) CreateCustomer(c echo.Context) (err error) {
	u := model.Custommer{}
	c.Bind(&u)

	cust, err := controller.interact.Add(u)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, cust)
	return
}

// GET | get all customer
func (controller *CustomerController) GetAll(c echo.Context) (err error) {
	users, err := controller.interact.CustAll()
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, users)
	return
}

// GET | by id
func (controller *CustomerController) GetById(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.interact.CustId(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)
	return
}

// PUT | update
func (controller *CustomerController) UpdateCust(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := model.Custommer{ID: id}

	user, err = controller.interact.Update(user)
	if err != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(201, user)
	return
}

// DELETE | delete
func (controller *CustomerController) DeleteCust(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := model.Custommer{ID: id}

	err = controller.interact.Delete(user)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)
	return
}
