package routes

import (
	"k6/token_db/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Default(router *gin.Engine) {
	router.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Server is Healthy"})
	})
}

func DemoRoute(router *gin.Engine, controller controllers.DemoController) {
	router.POST("/create", controller.CreateToken)
	router.POST("/store", controller.StoreData)
}
