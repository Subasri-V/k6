package controllers

import (
	"fmt"
	"k6/eg2/interfaces"
	"k6/eg2/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type DemoController struct {
	DemoService interfaces.IDemo
}

func InitDemoController(DemoService interfaces.IDemo) DemoController {
	return DemoController{DemoService}
}

func(d * DemoController) CreateToken(ctx *gin.Context) {
	var sample *models.Sample
	if err := ctx.ShouldBindJSON(&sample); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := d.DemoService.CreateToken(sample)
	fmt.Println(sample.Username)
	fmt.Println(sample.Password)
	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": result})
	ctx.JSON(http.StatusOK, gin.H{"Username": sample.Username, "Password": sample.Password})
}
