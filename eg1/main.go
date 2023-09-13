package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Sample struct {
	Text string `json:"text"`
}

func demoSample(ctx *gin.Context) {
	var sample Sample
	if err := ctx.ShouldBindJSON(&sample); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(sample.Text)
	ctx.JSON(http.StatusOK, gin.H{"Text": sample.Text})
}

func main() {
	s := gin.Default()
	s.POST("/demo", demoSample)
	s.Run(":8080")
}
