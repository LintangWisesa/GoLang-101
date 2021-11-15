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

func validate(token string) bool {
	
	claims := &JwtClaim{}

	tkn, err := jwt.ParseWithClaims(
		token,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return JWT_SIGNATURE_KEY, nil 
		},
	)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			// fmt.Println("Invalid signature")
			return false
		}
	}

	claims, ok := tkn.Claims.(*JwtClaim)
	if !ok {
		// fmt.Println("couldn't parse claims")
		return false
	}
	if claims.ExpiresAt < time.Now().UTC().Unix() {
		// fmt.Println("jwt is expired")
		return false
	}

	// fmt.Println(claims)
	// fmt.Println(claims.Mdn)
	// fmt.Println(time.Now().UTC().Unix())
	// fmt.Println(claims.ExpiresAt)
	// fmt.Println(claims.Issuer)
	// return claims.Mdn
	return true
}

func main(){
	t := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzY5NjMzODMsImlzcyI6Ik15U21hcnRmcmVuIiwibWRuIjoiMTIzNDU2YWJjZGUifQ.VXfhJJrpYurjUzQ4ZGFiS0jRqf1fuBa-xHSBVFoc4Qbr0uJuOXuVkH86_V8D2ctc05hbwUcLA1xAsr9C9nWblw"
	fmt.Println(validate(t))
}