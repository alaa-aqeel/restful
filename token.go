/*
	Package restful implements a simple library for restful-api with fiber.
	
*/
package restful 


import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"os"
	"errors"
)

// Create Token JWT 
// parameter : 
// 		id string : id user 
//      exp number : ExpiresAt By Minute (60 * 24) = 1 day 
// return token , error 
func (ctrl *Controller) CreateToken(id string, exp time.Duration) (string, error) {

	SecretKey := os.Getenv("TOEKN_SECRET_KEY")
	claims := &jwt.StandardClaims{}

	if exp != 0{

		claims.ExpiresAt= time.Now().Add(exp * time.Minute).Unix()
	}else{

		claims.ExpiresAt= time.Now().Add( ( 24 * 7 ) * time.Hour).Unix() // expires after 7 Day
	}

	if SecretKey == ""{

		return "", errors.New("Must provide a secret key under env variable TOEKN_SECRET_KEY")
	}

	if id != "" {
		claims.Id = id

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		
		return token.SignedString( []byte(SecretKey) )
	}

	return "", errors.New("The field id is empty")
}


// Check Token by tokenString 
// return Token, Clamis, error 
func VerifyToken(tokenString string ) (*jwt.Token, *jwt.StandardClaims, error) {

	SecretKey := os.Getenv("TOEKN_SECRET_KEY")

	claims := &jwt.StandardClaims{}

	token, error := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		return []byte(SecretKey), nil
	})


	return token, claims, error
}