package controllers

import (
	"fmt"
	"net/http"
	"transaction1to10/models"
	"transaction1to10/services"

	"github.com/gin-gonic/gin"
)

type TrasnactionController struct {
	transactionservice services.Transactionservice
}

func InitTrasnsactionController(transactionservice services.Transactionservice) TrasnactionController {
	return TrasnactionController{transactionservice}
}

func (t *TrasnactionController) MakeTransaction(ctx *gin.Context) {
	var transaction models.Transaction
	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		fmt.Println("Transaction binding problem: 🚨🚨🚨 ", err)
	}
	res, err1 := t.transactionservice.MakeTransaction(&transaction)
	if err1 != nil {
		fmt.Println("Controller to Service calling problem : 🚨🚨🚨", err1)
	}
	fmt.Println(res)
	ctx.JSON(http.StatusCreated, gin.H{"Transaction done 💳:  ": transaction.Ammount, "Grettings": "Transaction done 🙏🏼"})
}
