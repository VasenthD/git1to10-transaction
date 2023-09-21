package services

import (
	"context"
	"fmt"
	"transaction1to10/config"
	info "transaction1to10/infos"
	"transaction1to10/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Transactionservice struct {
	TCollection *mongo.Collection
	Tctx        context.Context
}

func InitTransactionService(collection *mongo.Collection, ctx context.Context) Transactionservice {
	return Transactionservice{collection, ctx}
}

func (s Transactionservice) MakeTransaction(transaction *models.Transaction) (string, error) {

	Ccollection := config.Getcollection(info.Dbname, "cusotmer")

//decreasing the ammount from the form_id user:
	fmt.Println(transaction.Ammount,"🎉 --- 🎉")
	_,err:=Ccollection.UpdateOne(s.Tctx, bson.M{"customer_id": transaction.From_id},
		bson.M{"$inc": bson.M{"balance": -(transaction.Ammount)}})
	if err != nil {
		fmt.Println("from user update problem : 🗿..🪦..🗿")
	}
	
//Increasing the ammount form the to_id user:

	_,err1 :=Ccollection.UpdateOne(s.Tctx, bson.M{"customer_id": transaction.To_id},
		bson.M{"$inc": bson.M{"balance": (transaction.Ammount)}})
		if err1 != nil {
			fmt.Println("from user update problem : 🗿..🪦..🗿")
		}
fmt.Println(transaction)
//Insergint the Transaction details in the sepreate database:
	res, err := s.TCollection.InsertOne(s.Tctx, transaction)
	if err != nil {
		fmt.Println("insert problem bro 🤖--> 💀 transaction file... 💀")
	}
	fmt.Println(res.InsertedID)

	return " 👍🏻 Inserted Sucessfully..💸 ", nil
}
