package controller

import (
	"mage_test_case/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (ac *ApiController) handleRegister(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		replyError(c, http.StatusBadRequest, err.Error())
		return
	}
	resp, err := ac.repo.Register(&req)
	if err != nil {
		replyError(c, http.StatusBadGateway, err.Error())
		return
	}
	reply(c, resp)
}

func (ac *ApiController) handleLogin(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		replyError(c, http.StatusBadRequest, err.Error())
		return
	}
	resp, err := ac.repo.Login(req)
	if err != nil {
		replyError(c, http.StatusBadGateway, err.Error())
		return
	}
	reply(c, resp)
}

func (ac *ApiController) handleEndGame(c *gin.Context) {
	var req model.EndGameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		replyError(c, http.StatusBadRequest, err.Error())
		return
	}
	resp, err := ac.repo.EndGame(req)
	if err != nil {
		replyError(c, http.StatusBadGateway, err.Error())
		return
	}
	reply(c, resp)
}

func (ac *ApiController) handleLeaderboard(c *gin.Context) {
	resp, err := ac.repo.GetLeaderboard()
	if err != nil {
		replyError(c, http.StatusBadGateway, err.Error())
		return
	}
	reply(c, resp)
}

func reply(c *gin.Context, data interface{}) {
	resp := model.SuccessResponse{Status: model.StatusSuccess, Timestamp: time.Now().Unix(), Result: data}
	c.JSON(http.StatusOK, resp)
}

func replyError(c *gin.Context, httpStatus int, err string) {
	c.JSON(httpStatus, model.FailResponse{Error: err, Status: model.StatusFail})
}
