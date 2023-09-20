package services

import (
	"context"
	"fmt"
	"transaction1to10/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Customerservice struct {
	Ccollection *mongo.Collection
	ctx         context.Context
}

func InitCustomerservice(collection *mongo.Collection, ctx context.Context) Customerservice {
	return Customerservice{collection, ctx}
}

func (s *Customerservice) CreateCustomer(Customer *models.Customer) (*models.DBresponse, error) {
	res, err := s.Ccollection.InsertOne(s.ctx, Customer)
	if err != nil {
		fmt.Println("services ERROR:", err)
		return nil, err
	}

	output := &models.DBresponse{
		ID:          res.InsertedID.(primitive.ObjectID).Hex(),
		Name:        Customer.Name,
		Customer_ID: Customer.Customer_ID,
	}

	fmt.Println(res, " New ID")
	return output, nil
}
