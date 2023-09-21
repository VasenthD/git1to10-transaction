package interfaces

import "transaction1to10/models"

type Customer interface {
	CreateCustomer(models.Customer) (models.DBresponse, error)
}
