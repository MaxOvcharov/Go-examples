package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var mySigningString = []byte("mysupersecretkey")


func GenerateJWT()(string, error){

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["autorized"] = true
	claims["user"] = "Max Ovcharov"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningString)

	if err != nil {
		fmt.Errorf("something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func main() {
	fmt.Println("Simple JWT client \n")

	tokenString, err := GenerateJWT()

	fmt.Println("JWT TOKEN: %s", tokenString)

	if err != nil {
		fmt.Println("Error generating token string")
	}
}