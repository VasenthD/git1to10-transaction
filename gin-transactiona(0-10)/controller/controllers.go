package controllers

import (
	"fmt"
	"net/http"
	"transaction1to10/models"
	"transaction1to10/services"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	cutomerservice services.Customerservice
}

func InitCustomerController(customerservices services.Customerservice) CustomerController {
	return CustomerController{customerservices}
}

func (c *CustomerController) CreateCustomer(ctx *gin.Context) {
	var customer *models.Customer

	if err := ctx.ShouldBindJSON(&customer); err != nil {
		fmt.Println("Error : it is not binding the json")
		ctx.JSON(http.StatusBadRequest, "Binding problem - check it 💀 ☠️!!!")
	}
	newCustomer, err := c.cutomerservice.CreateCustomer(customer)
	if err != nil {
		fmt.Println("Error : Cant call the services - check it !!!")
		ctx.JSON(http.StatusInternalServerError, "Calling the services is the problem 💀 ☠️!!!")
	}
	ctx.JSON(http.StatusCreated, gin.H{"Customer Created": newCustomer.ID, "Grettings": "Vanakam da mapla 🙏🏼😀🙏🏼", "🕶": "😈 welcome to greate Kirigalan magic show 👽"})
}
