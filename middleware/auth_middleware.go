package middleware

import (
	"net/http"

	"github.com/rizalbachtiar/belajar-golang-restful-api/helper"
	"github.com/rizalbachtiar/belajar-golang-restful-api/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "x1234" == request.Header.Get("X-API-Key") {
		//ok
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		//error
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAOUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}
