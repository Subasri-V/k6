package routes

import (
	"k6/eg2/controllers"

	"github.com/gin-gonic/gin"
)

func DemoRoute(router *gin.Engine, controller controllers.DemoController){
	router.POST("/create",controller.CreateToken)
}
