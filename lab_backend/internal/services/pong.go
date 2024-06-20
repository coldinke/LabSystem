package services

import (
	"github.com/gin-gonic/gin"
	"lab_backend/internal/handler"
)

func Pong(ctx *gin.Context) {
	handler.HandleResponse(ctx, nil, "Pong....")
}
