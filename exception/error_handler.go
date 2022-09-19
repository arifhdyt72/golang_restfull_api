package exception

import (
	"golang_restfull_api/helper"
	"golang_restfull_api/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writter http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writter, request, err) {
		return
	}

	if validationErrors(writter, request, err) {
		return
	}

	internalServerError(writter, request, err)
}

func validationErrors(writter http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writter.Header().Set("Content-Type", "application/json")
		writter.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(writter, webResponse)

		return true
	} else {
		return false
	}
}

func notFoundError(writter http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writter.Header().Set("Content-Type", "application/json")
		writter.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writter, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writter http.ResponseWriter, request *http.Request, err interface{}) {
	writter.Header().Set("Content-Type", "application/json")
	writter.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteToResponseBody(writter, webResponse)
}
