package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Respond(ctx *gin.Context, httpStatus int, body interface{}) {
	switch ctx.GetHeader("Accept") {
	case "application/xml":
		ctx.XML(httpStatus, body)
	default:
		ctx.JSON(httpStatus, body)
	}
}

func RespondOK(ctx *gin.Context, body interface{}) {
	Respond(ctx, http.StatusOK, body)
}

func RespondError(ctx *gin.Context, err *ApplicationError) {
	Respond(ctx, err.Status, err)
}
