package handles

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/noilpa/rest/utils"
	"net/http"
	"strings"
	"time"
)

var ErrTokenMissing = errors.New("Missing authorization token")
var ErrTokenInvalid = errors.New("Token is invalid")

// secret key
var key = []byte("a88bf3ca28776f7c2b1c27b7eea1bbf6")

// todo: referesh token

type JWTClaims struct {
	Usr string
	jwt.StandardClaims
}

func Decode(tokenStr string) (*JWTClaims, error){

	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	claims, ok := token.Claims.(*JWTClaims)

	if ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func Encode(usr utils.User) (string, error) {

	expireToken := time.Now().Add(time.Hour*24).Unix()

	claims:= JWTClaims{
		usr.Login,
		jwt.StandardClaims{
			ExpiresAt:expireToken,
			Issuer:"restream",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(key)
}


func JwtAuth(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			http.Error(w, ErrTokenMissing.Error(), 400)
			return
		}

		token, err := getToken(tokenHeader)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		claims, err := Decode(token)
		if err != nil {
			if err != nil {
				http.Error(w, ErrTokenInvalid.Error(), 400)
				return
			}
		}

		fmt.Printf("User %s enter with token\n", claims.Usr)

		h.ServeHTTP(w, r)

	})
}

func getToken(tokenHeader string) (string, error) {
	s := strings.Split(tokenHeader, " ")
	if len(s) < 2 {
		return "", ErrTokenMissing
	}
	return s[1], nil
}