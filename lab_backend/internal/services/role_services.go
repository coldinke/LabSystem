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

// GetRolesList
// @Summary get the all roles
// @Schemes
// @Description Retrieve all roles from the database and return the data sting as JSON
// @Tags Role
// @Accept json
// @Produce json
// @Success 200 {string} RoleList
// @Router /admin/api/roles [get]
func GetRolesList(ctx *gin.Context) {
	data, err := models.GetAllRoles(models.DB)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
	}
	handler.HandleResponse(ctx, nil, data)
}

type roleRequest struct {
	RoleType int            `json:"role_type"`
	Rights   []models.Right `json:"rights"`
}

// UpdateRole
// @Summary update role by ID
// @Schemes
// @Description update the role by ID.
// @Tags Role
// @Accept json
// @Produce json
// @Success 200 {string} success
// @Router /admin/api/roles/{id} [put]
func UpdateRole(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		myErr := errors.New(http.StatusBadRequest, reason.RequestFormatError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}

	var req roleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		myErr := errors.New(http.StatusBadRequest, reason.RequestFormatError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}

	data := map[string]interface{}{
		"RoleType": req.RoleType,
		"Rights":   req.Rights,
	}

	if err := models.UpdateRoleByID(models.DB, id, data); err != nil {
		myErr := errors.New(http.StatusConflict, reason.DatabaseError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}
	handler.HandleResponse(ctx, nil, reason.Success)
}

// DeleteRole
// @Summary Delete role by ID
// @Param id path int true "RoleID"
// @Schemes
// @Description Delete role By ID
// @Tags Role
// @Accept json
// @Produce json
// @Success 200 {string} RightList
// @Router /admin/api/roles/{id} [delete]
func DeleteRole(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		myErr := errors.New(http.StatusBadRequest, reason.RequestFormatError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}
	if _, err := models.GetRoleByID(models.DB, id); err != nil {
		myErr := errors.New(http.StatusNotFound, reason.UnknownError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}

	if err := models.DeleteRoleByID(models.DB, id); err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}

	handler.HandleResponse(ctx, nil, reason.Success)
}
