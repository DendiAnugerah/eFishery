package main

import (
	"net/http"

	"echo-crud/controller"

	"github.com/labstack/echo/v4"
)

// Bikin CRUD (tambahin Update, Delete, dan Get by id) + LOG Semua aktivitas pada endpointnya
// Kalau bisa pake loggingnya pake Logrus
// Export postman collection yang sudah dibuat

func main() {
	// initialize echo framework
	e := echo.New()

	// root router
	e.GET("/", func(c echo.Context) error {
		result := map[string]string{
			"response_code": "200",
			"message":       "Success",
		}

		return c.JSON(http.StatusOK, result)
	})

	cust := e.Group("customer")

	// Get all cust
	cust.GET("/", controller.GetCust)

	// Get cust by id
	cust.GET("/:id", controller.GetCustbyId)

	// Create cust
	cust.POST("/", controller.CreateCustomer)

	// Update cust
	cust.PUT("/:id", controller.UpdateCust)

	// Detel cust
	cust.DELETE("/:id", controller.DeleteCust)

	// start service secho
	e.Logger.Fatal(e.Start(":9090"))
}
