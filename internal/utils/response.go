package utils

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type Meta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}


func SuccessResponse(c *gin.Context, status int, msg string, data interface{}) {
	c.JSON(status, APIResponse{
		Success: true,
		Message: msg,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, status int, msg string) {
	c.JSON(status, APIResponse{
		Success: false,
		Message: msg,
	})
}

func ValidationErrorResponse(c *gin.Context, errs map[string]string) {
	c.JSON(http.StatusUnprocessableEntity, APIResponse{
		Success: false,
		Message: "Validation failed",
		Error:   errs,
	})
}

func PaginatedResponse(c *gin.Context, data interface{}, total int64, page, limit int) {
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Data retrieved",
		Data:    data,
		Meta: &Meta{
			Page:       page,
			Limit:      limit,
			Total:      total,
			TotalPages: totalPages,
		},
	})
}


func HandleServiceError(c *gin.Context, err error) {
	if appErr, ok := err.(*AppError); ok {
		c.JSON(appErr.StatusCode, APIResponse{
			Success: false,
			Message: appErr.Message,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, APIResponse{
		Success: false,
		Message: "Internal server error",
	})
}