package routes

import (
	"transaction1to10/controller"

	"github.com/gin-gonic/gin"
)

func Allroutes(routes *gin.Engine, controller controllers.CustomerController) {
	routes.POST("/create", controller.CreateCustomer)
}
