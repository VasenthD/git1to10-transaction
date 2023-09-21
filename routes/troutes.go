package routes

import (
	"transaction1to10/controller"

	"github.com/gin-gonic/gin"
)

func Alltroutes(routes *gin.Engine, controllers controllers.TrasnactionController){

	routes.POST("/transaction", controllers.MakeTransaction)

}
