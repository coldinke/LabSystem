package services

import (
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/errors"
	"lab_backend/internal/handler"
	"lab_backend/internal/models"
	"lab_backend/internal/reason"
	"net/http"
	"strconv"
)

func GetSensorList(ctx *gin.Context) {
	labID, err := strconv.ParseUint(ctx.Param("labID"), 10, 32)
	if err != nil {
		myErr := errors.New(http.StatusBadRequest, reason.RequestFormatError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}

	data, err := models.GetSensorList(models.DB, labID)

	if err != nil {
		handler.HandleResponse(ctx, err, nil)
	}
	handler.HandleResponse(ctx, nil, data)
}

func GetLastSensor(ctx *gin.Context) {
	labID, err := strconv.ParseUint(ctx.Param("labID"), 10, 32)
	if err != nil {
		myErr := errors.New(http.StatusBadRequest, reason.RequestFormatError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}

	data, err := models.GetLatestSensorDataByLabID(models.DB, uint(labID))

	if err != nil {
		handler.HandleResponse(ctx, err, nil)
	}
	handler.HandleResponse(ctx, nil, data)
}

func PublishMessage(ctx *gin.Context) {

}
