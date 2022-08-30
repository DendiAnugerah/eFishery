package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DendiA/crud/pkg/config"
	"github.com/DendiA/crud/pkg/controller"
	"github.com/labstack/echo/v4"
)

// Mas ini baru bisa create sama delete aja, sisanya besok, saya bingung, nanti saya coba pelajari lagi

func main() {
	// initialize echo framework
	e := echo.New()

	// connection to DB
	DB, _ := config.Conection()
	h := controller.New(DB)

	_, err := config.Conection()
	if err != nil {
		log.Panic(err)
		return
	}
	fmt.Println("Connected")

	// root router
	e.GET("/", func(c echo.Context) error {
		result := map[string]string{
			"response_code": "200",
			"message":       "Success",
		}

		return c.JSON(http.StatusOK, result)
	})

	cust := e.Group("customer")

	cust.POST("/", h.CreateCustomer)
	cust.GET("/", h.GetCust)
	cust.GET("/:id", h.GetCustbyId)
	cust.PUT("/:id", h.UpdateCust)
	cust.DELETE("/:id", h.DeleteCust)

	// // start service secho
	e.Logger.Fatal(e.Start(":9090"))
}
