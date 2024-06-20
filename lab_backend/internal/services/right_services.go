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

// GetRightsList
// @Summary get the all rights
// @Schemes
// @Description  Retrieve all rights from the database and return the data string as JSON
// @Tags Rights
// @Accept json
// @Produce json
// @Success 200 {string} RightList
// @Router /admin/api/rights [get]
func GetRightsList(ctx *gin.Context) {
	data, err := models.GetMenuList()
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
	}
	handler.HandleResponse(ctx, err, data)
}

type rightRequest struct {
	Title string `json:"title"`
	Path  string `json:"path"`
	Icon  string `json:"icon"`
}

// UpdateRight
// @Summary update the right by id
// @Schemes
// @Description  Update the right by id to database
// @Tags Rights
// @Success 200 {string} json "{"code": "200", "err":"", "data":""}"
// @Router /admin/api/rights/{id} [put]
func UpdateRight(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		myErr := errors.New(http.StatusBadRequest, reason.RequestFormatError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}

	var req rightRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		myErr := errors.New(http.StatusBadRequest, reason.RequestFormatError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}

	data := map[string]interface{}{
		"Title": req.Title,
		"Path":  req.Path,
		"Icon":  req.Icon,
	}

	if err := models.UpdateRightByID(models.DB, id, data); err != nil {
		myErr := errors.New(http.StatusConflict, reason.DatabaseError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}

	handler.HandleResponse(ctx, nil, reason.Success)
}

func DeleteRight(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		myErr := errors.New(http.StatusBadRequest, reason.RequestFormatError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}
	if _, err := models.GetRightByID(models.DB, id); err != nil {
		myErr := errors.New(http.StatusNotFound, reason.UnknownError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}

	if err := models.DeleteRightByID(models.DB, id); err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}

	handler.HandleResponse(ctx, nil, reason.Success)
}
