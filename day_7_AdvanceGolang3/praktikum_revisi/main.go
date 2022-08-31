package main

import (
	"custommer-crud/config"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	// initialize echo framework
	e := echo.New()

	// connection to DB
	DB, _ := config.Conection()
	//h := controller.New(DB)

	_, err := config.Conection()
	if err != nil {
		log.Panic(err)
		return
	}
	fmt.Println("Connected")

	cust := e.Group("customer")

	cust.POST("/", h.CreateCustomer)
	cust.GET("/", h.GetCust)
	cust.GET("/:id", h.GetCustbyId)
	cust.PUT("/:id", h.UpdateCust)
	cust.DELETE("/:id", h.DeleteCust)

	// // start service secho
	e.Logger.Fatal(e.Start(":9090"))
}
