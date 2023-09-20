package interfaces

import "transaction1to10/models"

type ITransaction interface{
	MakeTransaction(transactiong models.Transaction)(string,error)
}