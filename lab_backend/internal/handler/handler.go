package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	myErrors "github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
	"lab_backend/internal/reason"
	"net/http"
)

// HandleResponse Handle response body
func HandleResponse(ctx *gin.Context, err error, data interface{}) {
	if err == nil {
		ctx.JSON(http.StatusOK, NewRespBodyData(http.StatusOK, reason.Success, data))
		return
	}

	var myErr *myErrors.Error
	if !errors.As(err, &myErr) {
		log.Error(err, "\n", myErrors.LogStack(2, 5))
		ctx.JSON(http.StatusInternalServerError, NewRespBody(
			http.StatusInternalServerError, reason.UnknownError))
		return
	}

	if myErrors.IsInternalServer(myErr) {
		log.Error(myErr)
	}

	respBody := NewRespBodyFromError(myErr)
	if data != nil {
		respBody.Data = data
	}
	ctx.JSON(myErr.Code, respBody)
}
