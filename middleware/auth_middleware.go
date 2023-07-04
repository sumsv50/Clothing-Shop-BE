package middleware

import (
	"clothing-shop/handler"
	"clothing-shop/service"
	"net/http"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("Authorization")
		if token == "" {
			handler.JSON(w, http.StatusUnauthorized, "", nil)
			return
		}
		err := service.ValidateToken(token)
		if err != nil {
			handler.JSON(w, http.StatusUnauthorized, "", nil)
			return
		}

		next.ServeHTTP(w, req)
	})
}
