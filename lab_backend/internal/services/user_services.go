package services

import (
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/errors"
	"lab_backend/internal/handler"
	"lab_backend/internal/middleware"
	"lab_backend/internal/models"
	"lab_backend/internal/reason"
	"net/http"
	"strconv"
)

// GetUsersList
// @Summary get the all users
// @Schemes
// @Description  Retrieve all users from the database and return the data string as JSON
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {string} Users
// @Router /admin/api/users [get]
func GetUsersList(ctx *gin.Context) {
	data, err := models.GetUserList(models.DB)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	handler.HandleResponse(ctx, err, data)
}

func GetUsersCount(ctx *gin.Context) {
	var count int64
	err := models.DB.Model(models.User{}).Count(&count).Error
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	handler.HandleResponse(ctx, nil, count)
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login
func Login(ctx *gin.Context) {
	var req LoginRequest

	if err := ctx.BindJSON(&req); err != nil {
		handler.HandleResponse(ctx, errors.New(http.StatusBadRequest, reason.ForbiddenError), nil)
		return
	}

	username := req.Username
	password := req.Password
	if username == "" || password == "" {
		return
	}
	user, err := models.CheckUser(models.DB, username, password)
	if err != nil {
		handler.HandleResponse(ctx, errors.New(http.StatusForbidden, reason.UnauthorizedError), nil)
		return
	}

	token, err := middleware.GenerateToken(username)
	if err != nil {
		handler.HandleResponse(ctx, errors.New(http.StatusInternalServerError, reason.UnauthorizedError), err)
		return
	}

	response := map[string]interface{}{
		"user":  user,
		"token": token,
	}
	handler.HandleResponse(ctx, nil, response)
}

type userRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   int    `json:"roleID"`
}

func UpdateUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		myErr := errors.New(http.StatusBadRequest, reason.RequestFormatError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}

	var req userRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		myErr := errors.New(http.StatusBadRequest, reason.RequestFormatError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}

	data := map[string]interface{}{
		"username": req.Username,
		"password": req.Password,
		"roleID":   req.RoleID,
	}

	if err := models.UpdateUserByID(models.DB, id, data); err != nil {
		myErr := errors.New(http.StatusConflict, reason.DatabaseError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}
	handler.HandleResponse(ctx, nil, reason.Success)
}

func NewUser(ctx *gin.Context) {
	var req userRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		myErr := errors.New(http.StatusBadRequest, reason.RequestFormatError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}

	data := map[string]interface{}{
		"username": req.Username,
		"password": req.Password,
		"roleID":   req.RoleID,
	}

	if err := models.CreateUser(models.DB, data); err != nil {
		myErr := errors.New(http.StatusConflict, reason.DatabaseError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}
	handler.HandleResponse(ctx, nil, reason.Success)
}

func DeleteUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		myErr := errors.New(http.StatusBadRequest, reason.RequestFormatError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}
	if _, err := models.GetUserByID(models.DB, id); err != nil {
		myErr := errors.New(http.StatusNotFound, reason.UnknownError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}

	if err := models.DeleteUserByID(models.DB, id); err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}

	handler.HandleResponse(ctx, nil, reason.Success)
}
