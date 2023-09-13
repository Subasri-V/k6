package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	// // Define your claims (payload)
	// claims := jwt.MapClaims{
	//     "sub": "1234567890",
	//     "name": "subasri",
	//     "iat": time.Now().Unix(),
	// }

	// // Define your secret key (keep it secret!)
	// secretKey := []byte("your-secret-key")

	// // Create a new token object with claims
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// // Sign the token with the secret key to get the final token as a string
	// tokenString, err := token.SignedString(secretKey)
	// if err != nil {
	//     fmt.Println(err)
	//     return
	// }

	// fmt.Println(tokenString) // This is the JWT token

	// The JWT token you want to decode and verify
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2OTQ1ODEwNjQsIm5hbWUiOiJzdWJhc3JpIiwic3ViIjoiMTIzNDU2Nzg5MCJ9.INtEnZmdW5s4Du3KDZgV3TU4uV_dmlCxtZ_2wrc74y8"

	// Define your secret key (must be the same as the one used for signing)
	secretKey := []byte("your-secret-key")

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
		return
	}

	// Check if the token is valid
	if token.Valid {
		// Access the claims from the payload
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("Invalid claims format")
			return
		}

		// Access specific claim values
		subject := claims["sub"].(string)
		name := claims["name"].(string)

		fmt.Printf("Subject: %s\n", subject)
		fmt.Printf("Name: %s\n", name)
	} else {
		fmt.Println("Token is not valid")
	}

}
