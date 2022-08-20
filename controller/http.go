package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (gc *ApiController) handleRegister(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (gc *ApiController) handleLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (gc *ApiController) handleEndGame(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (gc *ApiController) handleLeaderboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
