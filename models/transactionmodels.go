package models

type Transaction struct{
	Ammount int64 `json:"ammount" bson:"ammount"`
	From_id string `json:"from_id" bson:"from_id"`
	To_id string `json:"to_id" bson:"to_id"`
}

