package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func WithJWTAuth(handlerfunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get the token from the request header
		tokenString := getTokenFromRequest(r)

		//validate the token
		token, err := validateJWT(tokenString)
		if err != nil {
			log.Println("failed to authenticate token")
			permissionDenied(w)
		}

		if !token.Valid {
			log.Println("failed to authenticate token")
			permissionDenied(w)
		}
		//get the userID from the token
		//call the handler func and continue the endpoint
	}
}

func permissionDenied(w http.ResponseWriter) {
	WriteJSON(w, http.StatusUnauthorized, ErrorResponse{
		Error: fmt.Errorf("Permission denied").Error(),
	})
	return
}

func getTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	tokenQuery := r.URL.Query().Get("token")

	if tokenAuth != "" {
		return tokenAuth
	}

	if tokenQuery != "" {
		return tokenQuery
	}

	return ""
}

func validateJWT(t string) (*jwt.Token, error) {
	secret := Envs.JWTSecret

	return jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}

		return []byte(secret), nil
	})
}
