package users

import (
	"github.com/labstack/echo"
	"money-account/memoryDB"
	"net/http"
)

func Accounts(c echo.Context) error {
	accounts := memoryDB.GetAllAccounts()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"accounts":accounts,
	})
}
