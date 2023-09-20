package main

import (
	"context"
	"fmt"
	"log"
	"transaction1to10/config"
	controllers "transaction1to10/controller"
	info "transaction1to10/infos"
	"transaction1to10/routes"
	"transaction1to10/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoclient *mongo.Client
	ctx         context.Context
	server      *gin.Engine
)

func initRoutes() {
	routes.Default(server)
}

func CinitApp(mongoClient *mongo.Client) {
	ctx = context.TODO()
	// creating a collection
	profileCollection := mongoClient.Database(info.Dbname).Collection("cusotmer")
	//pass the collection and context to instances the service
	profileService := services.InitCustomerservice(profileCollection, ctx)
	//pass the service to instantieate the controller
	profileController := controllers.InitCustomerController(profileService)
	//pass the controller to route
	routes.Allroutes(server, profileController)
}

func TinitApp(mongoClient *mongo.Client) {
	ctx = context.TODO()
	// creating a collection
	TCollection := mongoClient.Database(info.Dbname).Collection("Transaction")
	//pass the collection and context to instances the service
	TService := services.InitCustomerservice(TCollection, ctx)
	//pass the service to instantieate the controller
	TController := controllers.InitCustomerController(TService)
	//pass the controller to route
	routes.Allroutes(server, TController)
}

func main() {
	server = gin.Default()
	mongoclient, err := config.ConnectDatabase()
	defer mongoclient.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
	initRoutes()
	CinitApp(mongoclient)
	TinitApp(mongoclient)
	fmt.Println("server running on port", info.Port)
	log.Fatal(server.Run(info.Port))
}
