package router

import (
	"log"
	"net/http"
	"strings"
)

//
//func checkPlatformHeader(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		platform := r.Header.Get("platform")
//		log.Println(r.RequestURI)
//		if platform == "website" {
//			//next.ServeHTTP(w, r)
//			checkAuthorization(next, w, r)
//		} else {
//			http.Error(w, "Unauthorized", http.StatusUnauthorized)
//			return
//		}
//	})
//}

func CheckAuthorization(next http.Handler) http.Handler  {
	// This function only checks if authorization header contains token in "Bearer {{token}}
	// This does not actually check if token is valid
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		authHeader := r.Header.Get("Authorization")
		tokenArr := strings.Split(authHeader, " ")
		if len(tokenArr) > 1 {
			token := tokenArr[1]

			if token == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	})
}
