package services

import (
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/errors"
	"lab_backend/internal/handler"
	"lab_backend/internal/models"
	"lab_backend/internal/reason"
	"lab_backend/internal/serializer"
	"net/http"
	"strconv"
	"time"
)

// GetLabsList
// @Summary get the all labs
// @Schemes
// @Description  Retrieve all labs from the database and return the data string as JSON
// @Tags Count
// @Accept json
// @Produce json
// @Success 200 {string} LabList
// @Router /admin/api/labs [get]
func GetLabsList(ctx *gin.Context) {
	data, err := models.GetLabList(models.DB)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	handler.HandleResponse(ctx, nil, data)
}

// GetLabsCount
// @Summary count of the labs
// @Schemes
// @Description count of the labs
// @Tags Count
// @Accept json
// @Produce json
// @Success 200 {int} count
// @Router /admin/api/labs/count [get]
func GetLabsCount(ctx *gin.Context) {
	var count int64
	err := models.DB.Model(models.Lab{}).Count(&count).Error
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	handler.HandleResponse(ctx, nil, count)
}

// GetBooksCount
// @Summary count of the books
// @Schemes
// @Description  count of the books
// @Tags Count
// @Accept json
// @Produce json
// @Success 200 {int} count
// @Router /admin/api/books/count[get]
func GetBooksCount(ctx *gin.Context) {
	var count int64
	err := models.DB.Model(models.Book{}).Count(&count).Error
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	handler.HandleResponse(ctx, nil, count)
}

type labRequest struct {
	LabName     string `json:"lab_name"`
	LabType     int    `json:"lab_type"`
	RiskType    int    `json:"risk_type"`
	CollegeType int    `json:"college_type"`
	Capacity    int    `json:"capacity"`
}

func UpdateLab(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		myErr := errors.New(http.StatusBadRequest, reason.RequestFormatError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}

	var req labRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		myErr := errors.New(http.StatusBadRequest, reason.RequestFormatError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}

	data := map[string]interface{}{
		"lab_name":     req.LabName,
		"risk_type":    req.RiskType,
		"college_type": req.CollegeType,
		"lab_type":     req.LabType,
		"capacity":     req.Capacity,
	}

	if err := models.UpdateLabByID(models.DB, id, data); err != nil {
		myErr := errors.New(http.StatusConflict, reason.DatabaseError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}
	handler.HandleResponse(ctx, nil, reason.Success)
}

func NewLab(ctx *gin.Context) {
	var req labRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		myErr := errors.New(http.StatusBadRequest, reason.RequestFormatError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}

	data := map[string]interface{}{
		"lab_name":     req.LabName,
		"lab_type":     req.LabType,
		"risk_type":    req.RiskType,
		"college_type": req.CollegeType,
		"capacity":     req.Capacity,
	}

	if err := models.CreateLab(models.DB, data); err != nil {
		myErr := errors.New(http.StatusConflict, reason.DatabaseError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}
	handler.HandleResponse(ctx, nil, reason.Success)
}

func DeleteLab(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		myErr := errors.New(http.StatusBadRequest, reason.RequestFormatError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}
	if _, err := models.GetLabByID(models.DB, id); err != nil {
		myErr := errors.New(http.StatusNotFound, reason.UnknownError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}

	if err := models.DeleteLabByID(models.DB, id); err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}

	handler.HandleResponse(ctx, nil, reason.Success)
}

func GetLabDetail(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		myErr := errors.New(http.StatusBadRequest, reason.RequestFormatError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}

	lab, err := models.GetLabDetailByID(models.DB, id)
	if err != nil {
		myErr := errors.New(http.StatusConflict, reason.DatabaseError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}
	handler.HandleResponse(ctx, nil, lab)
}

func GetBookList(ctx *gin.Context) {
	data, err := models.GetBookList(models.DB)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	handler.HandleResponse(ctx, nil, data)
}

func NewBook(ctx *gin.Context) {
	var req serializer.BookingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		myErr := errors.New(http.StatusBadRequest, reason.RequestFormatError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}

	book := models.Book{
		ClassID:  req.BookClass,
		Reason:   req.BookReason,
		BookTime: req.BookTime,
		LabID:    uint(req.LabID),
		Username: req.BookUsername,
	}

	if err := models.CreateBook(models.DB, book); err != nil {
		myErr := errors.New(http.StatusConflict, reason.DatabaseError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}
	handler.HandleResponse(ctx, nil, reason.Success)
}

func GetBookListByUsername(ctx *gin.Context) {
	username := ctx.Param("username")

	data, err := models.GetBookListByUsername(models.DB, username)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	handler.HandleResponse(ctx, nil, data)
}

func GetBookListByLabID(ctx *gin.Context) {
	labID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		myErr := errors.New(http.StatusBadRequest, reason.RequestFormatError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}

	bookTime, err := time.Parse(time.RFC3339, ctx.Query("book_time"))
	if err != nil {
		myErr := errors.New(http.StatusAccepted, reason.RequestFormatError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}

	data, err := models.GetBookListByLabID(models.DB, labID, bookTime)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	handler.HandleResponse(ctx, nil, data)
}

func DeleteBook(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		myErr := errors.New(http.StatusBadRequest, reason.RequestFormatError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}
	if _, err := models.GetBookByID(models.DB, id); err != nil {
		myErr := errors.New(http.StatusNotFound, reason.UnknownError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}

	if err := models.DeleteBookByID(models.DB, id); err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}

	handler.HandleResponse(ctx, nil, reason.Success)
}

func GetBookListByStatus(ctx *gin.Context) {
	data, err := models.GetBookListByStatus(models.DB, 2)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	handler.HandleResponse(ctx, nil, data)
}

type updateBook struct {
	State int `json:"status"`
}

func UpdateBook(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		myErr := errors.New(http.StatusBadRequest, reason.RequestFormatError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}

	var req updateBook
	if err := ctx.ShouldBindJSON(&req); err != nil {
		myErr := errors.New(http.StatusBadRequest, reason.RequestFormatError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}

	if err := models.UpdateBookByID(models.DB, id, req.State); err != nil {
		myErr := errors.New(http.StatusConflict, reason.DatabaseError)
		handler.HandleResponse(ctx, myErr, nil)
		return
	}
	handler.HandleResponse(ctx, nil, reason.Success)
}
