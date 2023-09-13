package main

import (
	"context"
	"fmt"
	"k6/token_db/config"
	"k6/token_db/constants"
	"k6/token_db/controllers"
	"k6/token_db/services"
	"k6/token_db/routes"

	"log"

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

func initApp(mongoClient *mongo.Client) {
	ctx = context.TODO()
	demoCollection := mongoClient.Database(constants.DatabaseName).Collection("k6-token")
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
	initRoutes()
	initApp(mongoclient)
	fmt.Println("server running on port", constants.Port)
	log.Fatal(server.Run(constants.Port))
}
