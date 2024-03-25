package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type response struct {
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
}

// Success escribe una respuesta exitosa
func Success(ctx *gin.Context, status int, data interface{}, msg string) {
	ctx.JSON(status, response{
		Data:    data,
		Status:  status,
		Message: msg,
	})
}

// Failure escribe una respuesta fallida
func Failure(ctx *gin.Context, status int, err error) {
	ctx.JSON(status, errorResponse{
		Message: err.Error(),
		Status:  status,
		Code:    http.StatusText(status),
	})
}
