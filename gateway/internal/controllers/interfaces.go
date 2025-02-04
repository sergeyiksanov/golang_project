package controllers

import "github.com/gin-gonic/gin"

type IAuthController interface {
	SignIn(c *gin.Context)
	SignUp(c *gin.Context)
	RefreshTokens(c *gin.Context)
	Logout(c *gin.Context)
}
