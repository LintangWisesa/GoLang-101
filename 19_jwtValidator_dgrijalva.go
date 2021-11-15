
// "github.com/dgrijalva/jwt-go": deprecated!

package main

import (
	"fmt"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
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
	t := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzY5NjM5MTUsImlzcyI6Ik15U21hcnRmcmVuIiwibWRuIjoiMTIzNDU2YWJjZGUifQ.lN930x6AFKIvBrgcoIKxIV0YHo0xnaofHKBP2m_Q7eux7XokY2D1Aryt5NnRbPCocKUquoVYQ9Mm9_zqvoT-JQ"
	fmt.Println(validate(t))
}