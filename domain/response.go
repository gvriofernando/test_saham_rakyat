package domain

import "github.com/labstack/echo/v4"

type Response struct {
	ErrorCode int32       `json:"errorCode"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
}

func NewResponse(ctx echo.Context, statusCode int, message string, data interface{}) {
	ctx.JSON(statusCode, Response{
		ErrorCode: 0,
		Message:   message,
		Data:      data,
	})
}
