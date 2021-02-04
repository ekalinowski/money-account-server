package memoryDB

import (
	"fmt"
	"github.com/twinj/uuid"
	"money-account/entities"
	"net/http"
	"strings"
	"sync"
	"time"
)

var databaseTransactions *entities.HistoryDatabase
var databaseAccounts *entities.AccountsDatabase
var mutex = &sync.Mutex{}

func getAccountsDatabase() *entities.AccountsDatabase {
	if databaseAccounts == nil {
		databaseAccounts = &entities.AccountsDatabase{
			Id: uuid.NewV4().String(),
			User: "Admin",
			Value: 0.0,
		}
	}

	return databaseAccounts
}

func getHistoryDatabase() *entities.HistoryDatabase {
	if databaseTransactions == nil {
		databaseTransactions = &entities.HistoryDatabase{
			Id: uuid.NewV4().String(),
			Transactions: []entities.Transaction{

			},
		}
	}

	return databaseTransactions
}

func GetAllAccounts() *entities.AccountsDatabase {
	return getAccountsDatabase()
}

func GetAllTransactions() []entities.Transaction {
	return getHistoryDatabase().Transactions
}

func GetTransactionById(id string) *entities.Transaction {
	transactions := getHistoryDatabase()
	for _,transaction := range transactions.Transactions {
		if transaction.Id == id {
			return &transaction
		}
	}
	return nil
}

func UpdateAccount(transaction *entities.Transaction) (int,string){
	accounts := getAccountsDatabase()
	value := transaction.Amount
	if strings.ToUpper(transaction.Type) == "DEBIT" {
		if value > accounts.Value {
			return http.StatusUnauthorized, "Your account does not have enough balance for this operation"
		}
		value = value*-1
	}
	accounts.Value = accounts.Value + value
	id := saveTransaction(transaction)

	return http.StatusOK, "Successful operation, his id is: "+ id + ", your new balance is " + fmt.Sprintf("%3f", accounts.Value)
}

func saveTransaction(transaction *entities.Transaction) string {
	transaction.Id = uuid.NewV4().String()
	go storage(transaction)
	return transaction.Id

}

func storage(transaction *entities.Transaction) {
	transaction.EffectiveDate = time.Now().String()
	mutex.Lock()
	databaseTransactions.Transactions = append(getHistoryDatabase().Transactions, *transaction)
	mutex.Unlock()
}

