package users

import (
	"github.com/labstack/echo"
	"money-account/entities"
	"money-account/memoryDB"
	"net/http"
	"strings"
)

func MakeTransaction(c echo.Context) error {
	transaction := new(entities.Transaction)
	if err := c.Bind(transaction); err != nil {
		return err
	}

	if transaction.Type == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":"Type body param is necessary",
		})
	}

	if transaction.Amount <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":"Amount body param is necessary and must be positive",
		})
	}

	if strings.ToUpper(transaction.Type) != "CREDIT" && strings.ToUpper(transaction.Type) != "DEBIT"{
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Transaction type must be credit or debit",
		})
	}

	status, message := memoryDB.UpdateAccount(transaction)

	return c.JSON(status, map[string]interface{}{
		"message": message,
	})
}

func TransactionHistory(c echo.Context) error {
	id := c.QueryParam("id")

	if id == "" {
		transactions := memoryDB.GetAllTransactions()
		return c.JSON(http.StatusOK, map[string]interface{}{
			"transactions": transactions,
		})
	}


	transaction := memoryDB.GetTransactionById(id)
	if transaction == nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "This ID doesn't exist.",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"transaction": transaction,
	})
}


