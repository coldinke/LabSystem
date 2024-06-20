package handler

import (
	"github.com/segmentfault/pacman/errors"
)

// RespBody response Body
type RespBody struct {
	Code    int         `json:"code"`
	Reason  string      `json:"reason"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// NewRespBody new response Body
func NewRespBody(code int, reason string) *RespBody {
	return &RespBody{
		Code:   code,
		Reason: reason,
	}
}

// NewRespBodyFromError new response body from error
func NewRespBodyFromError(e *errors.Error) *RespBody {
	return &RespBody{
		Code:    e.Code,
		Reason:  e.Reason,
		Message: e.Message,
	}
}

// NewRespBodyData new response body with data
func NewRespBodyData(code int, reason string, data interface{}) *RespBody {
	return &RespBody{
		Code:   code,
		Reason: reason,
		Data:   data,
	}
}
