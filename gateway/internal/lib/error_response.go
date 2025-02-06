package lib

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Error string `json:"error,omitempty"`
}

func AbortErrorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, ErrorResponse{msg})
}
