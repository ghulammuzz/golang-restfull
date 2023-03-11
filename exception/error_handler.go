package exception

import (
	"ghulammuzz/golang-restfull/helper"
	"ghulammuzz/golang-restfull/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	
	if notFound(writer, request, err) {
		return 
	}

	if validationError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool{
	exception, ok := err.(validator.ValidationErrors)
	if ok{
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code: http.StatusBadRequest,
			Status: "Bad Reequest",
			Data : exception.Error(),
		}
		helper.WriteFromRequestBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func notFound(writer http.ResponseWriter, request *http.Request, err interface{}) bool{
	exception, ok := err.(NotFound)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code: http.StatusNotFound,
			Status: "Not Found",
			Data : exception.Error,
		}
		helper.WriteFromRequestBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code: http.StatusInternalServerError,
		Status: "Internal Status Error",
		Data : err,
	}
	helper.WriteFromRequestBody(writer, webResponse)

}