package main

import "net/http"

func WithJWTAuth(handlerfunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get the token from the request header
		tokenString := GetTokenFromRequest(r)

		//validate the token
		token, err := ValidateJWT(tokenString)

		//get the userID from the token
		//call the handler func and continue the endpoint
	}
}

func GetTokenFromRequest(r *http.Request) string {
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
