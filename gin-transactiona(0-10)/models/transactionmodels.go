package models

type Transaction struct{
	Ammount int64 `json:"ammount" bson:"ammount"`
	From_id int64 `json:"from_id" bson:"from_id"`
	To_id int64 `json:"to_id" bson:"to_id"`
}

