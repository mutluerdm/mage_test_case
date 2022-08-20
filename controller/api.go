package controller

import (
	"mage_test_case/config"
	"mage_test_case/mlog"
	"mage_test_case/repo"
	"sync"

	"github.com/gin-gonic/gin"
)

type ApiController struct {
	engine *gin.Engine
	repo   *repo.Repo
	mu     sync.RWMutex
}

func NewGame(cnf *config.Config) *ApiController {
	err, r := repo.New(cnf)
	if err != nil {
		mlog.PrintErrf("Repo Error : %+v", err)
		return nil
	}
	ac := ApiController{
		repo: r,
		mu:   sync.RWMutex{},
	}
	router := gin.Default()
	router.StaticFile("/home.html", "./home.html")
	router.POST("/register", ac.handleRegister)
	router.POST("/login", ac.handleLogin)
	router.POST("/endgame", ac.handleEndGame)
	router.POST("/leaderboard", ac.handleLeaderboard)
	ac.engine = router
	return &ac
}
