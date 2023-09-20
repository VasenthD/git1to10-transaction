package models

type Customer struct {
	Name        string  `json:"name" bson:"name"`
	Customer_ID string  `json:"customer_id" bson:"customer`
	BankId      string  `json:"bank_id" bson:"bank_id`
	Balance     float64 `json:"balance" bson:"balance`
}

type DBresponse struct {
	ID          string `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Customer_ID string `json:"customer_id`
}
