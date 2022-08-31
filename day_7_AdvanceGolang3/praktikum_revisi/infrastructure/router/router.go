package router

import (
	"custommer-crud/controller"
	"custommer-crud/infrastructure/config"

	"github.com/labstack/echo/v4"
)

func Init() {
	// initialize echo framework
	e := echo.New()

	cust := e.Group("customer")
	custControl := controller.ConnUser(config.Conection())

	cust.POST("/", func(c echo.Context) error { return custControl.CreateCustomer(c) })
	cust.GET("/", func(c echo.Context) error { return custControl.GetAll(c) })
	cust.GET("/:id", func(c echo.Context) error { return custControl.GetById(c) })
	cust.PUT("/:id", func(c echo.Context) error { return custControl.UpdateCust(c) })
	cust.DELETE("/:id", func(c echo.Context) error { return custControl.DeleteCust(c) })

	// // start service secho
	e.Logger.Fatal(e.Start(":9090"))
}
