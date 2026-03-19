package handler

import (
	"forma/pkg/response"
	"forma/pkg/token"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type AdminLoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func AdminLogin(c *gin.Context) {
	var req AdminLoginRequest

	// body parse
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c.Writer, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	// .env dan olish
	adminLogin := os.Getenv("ADMIN_LOGIN")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	// tekshirish
	if req.Login != adminLogin || req.Password != adminPassword {
		response.Error(c.Writer, http.StatusUnauthorized, "Invalid login or password", nil)
		return
	}

	// token yaratish
	tokenStr, err := token.GenerateAdminToken()
	if err != nil {
		response.Error(c.Writer, http.StatusInternalServerError, "Failed to generate token", nil)
		return
	}

	// success response (token bilan)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"token":   tokenStr,
	})
}