package main

import (
	"fmt"
	"time"
	jwt "github.com/golang-jwt/jwt"
)

type JwtClaim struct {
    jwt.StandardClaims
    Mdn string `json:"mdn"`
}

var JWT_SIGNING_METHOD = jwt.SigningMethodHS512
var JWT_SIGNATURE_KEY = []byte("321@nerftramSyM")
var JWT_ISSUER = "MySmartfren"
var JWT_MINUTES_EXP = 60

func generate(mdn string) string {
	
	claim := JwtClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    JWT_ISSUER,
			ExpiresAt: time.Now().Add(time.Duration(JWT_MINUTES_EXP) * time.Minute).Unix(),
		},
		Mdn: mdn,
	}
	
	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claim,
	)
	
	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return "error"
	}
	return signedToken
}

func main(){
	fmt.Println(generate("123456abcde"))
}