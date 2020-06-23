package restful 


import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"os"
	"errors"
)


func (ctrl *Controller) CreateToken(id string, exp time.Duration) (string, error) {

	SecretKey := os.Getenv("TOEKN_SECRET_KEY")
	claims := &jwt.StandardClaims{}

	if exp != 0{

		claims.ExpiresAt= time.Now().Add(exp * time.Minute).Unix()
	}else{

		claims.ExpiresAt= time.Now().Add( ( 24 * 7 ) * time.Hour).Unix() // expires after 7 Day
	}

	if id != "" && SecretKey != "" {
		claims.Id = id

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		
		return token.SignedString( []byte(SecretKey) )
	}

	return "", errors.New("Must provide a secret key under env variable TOEKN_SECRET_KEY")
}

func VerifyToken(tokenString string ) (*jwt.StandardClaims, error) {

	SecretKey := os.Getenv("TOEKN_SECRET_KEY")

	claims := &jwt.StandardClaims{}

	_, error := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		return []byte(SecretKey), nil
	})


	return claims, error

}


