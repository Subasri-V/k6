package controllers

import (
	"fmt"
	"k6/token_db/interfaces"
	"k6/token_db/models"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type DemoController struct {
	DemoService interfaces.IDemo
}

func InitDemoController(DemoService interfaces.IDemo) DemoController {
	return DemoController{DemoService}
}

func (d *DemoController) CreateToken(ctx *gin.Context) {
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

func (d *DemoController) StoreData(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		fmt.Println("error1")
		return
	}
	var secretKey = []byte("your-256-bit-secret")
	data, err := VerifyAndValidateToken(tokenString, secretKey)
	if err != nil {
		fmt.Println("error2")
		return
	}
	dbname := data["name"].(string)
	dbpwd := data["password"].(string)
	fmt.Println(dbname)

	sample := models.Sample{
		Username: dbname,
		Password: dbpwd,
	}
	result, err := d.DemoService.StoreData(&sample)
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

func VerifyAndValidateToken(tokenString string, secretKey []byte) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method (algorithm)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	fmt.Println(token)
	fmt.Printf("\n")

	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	// Check if the token is valid
	if token.Valid {
		// Access the claims from the payload
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("Invalid claims format")
			return nil, err
		}
		return claims, nil

		// // Access specific claim values
		// subject := claims["sub"].(string)
		// name := claims["name"].(string)

		// fmt.Printf("Subject: %s\n", subject)
		// fmt.Printf("Name: %s\n", name)
	} else {
		fmt.Println("Token is not valid")
		return nil,err
	}
}
