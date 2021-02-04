package entities

type AccountsDatabase struct {
	Id string `json:"id"`
	User string `json:"user"`
	Value float64 `json:"value"`
}

type HistoryDatabase struct {
	Id string `json:"id"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Amount float64 `json:"amount"`
	EffectiveDate string `json:"effectiveDate"`
}