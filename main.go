package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"money-account/users"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/healthcheck",  healthcheck)
	e.GET("/balance",  users.Accounts)
	e.POST("/transaction",  users.MakeTransaction)
	e.GET("/history",  users.TransactionHistory)


	e.Logger.Fatal(e.Start(":8080"))
}



func healthcheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":"Go money accounting system is running!",
	})
}