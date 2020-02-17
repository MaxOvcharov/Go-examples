package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

// Better put mySigningString into ENV
//mySigningString = os.Get("JWT_TOKEN_PHRASE")

var mySigningString = []byte("mysupersecretkey")

func HomePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, "JWT TOKEN: " + validToken)
}

func HandleRequests() {
	http.HandleFunc("/", HomePage)

	log.Fatal(http.ListenAndServe(":9001", nil))
}

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
	fmt.Println("Simple JWT client")

	HandleRequests()
}