package middleware

import (
	"golang_restfull_api/helper"
	"golang_restfull_api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writter http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("X-API-KEY") {
		// OK
		middleware.Handler.ServeHTTP(writter, request)
	} else {
		// error
		writter.Header().Set("Content-Type", "application/json")
		writter.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.WriteToResponseBody(writter, webResponse)
	}
}
