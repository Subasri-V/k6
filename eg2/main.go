package main

import (
	"context"
	"fmt"
	"k6/eg2/config"
	"k6/eg2/constants"
	"k6/eg2/controllers"
	"k6/eg2/routes"
	"k6/eg2/services"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


var (
	mongoclient *mongo.Client
	ctx         context.Context
	server      *gin.Engine
)
// func initRoutes() {
// 	routes.Default(server)
// }


func initApp(mongoClient *mongo.Client) {
	ctx = context.TODO()
	demoCollection := mongoClient.Database(constants.DatabaseName).Collection("k6-sample")
	demoService := services.InitializeCustomerService(ctx, demoCollection, mongoClient)
	demoController := controllers.InitDemoController(demoService)
	routes.DemoRoute(server, demoController)
}

func main() {
	server = gin.Default()
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
	//initRoutes()
	initApp(mongoclient)
	fmt.Println("server running on port", constants.Port)
	log.Fatal(server.Run(constants.Port))
}
