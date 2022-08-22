package controller

import (
	"context"
	"log"
	"mage_test_case/config"
	"mage_test_case/mlog"
	"mage_test_case/repo"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gin-gonic/gin"
)

type ApiController struct {
	engine *gin.Engine
	repo   *repo.Repo
	mu     sync.RWMutex
}

func NewAPI(cnf *config.Config) (*ApiController, error) {
	r, err := repo.New(cnf)
	if err != nil {
		mlog.PrintErrf("Repo Error : %+v", err)
		return nil, err
	}
	ac := ApiController{
		repo: r,
		mu:   sync.RWMutex{},
	}
	router := gin.Default()
	router.StaticFile("test.html", "./test.html")
	routerV1 := router.Group("/v1")
	routerV1.POST("user/register", ac.handleRegister)
	routerV1.POST("user/login", ac.handleLogin)
	routerV1.POST("endgame", ac.handleEndGame)
	routerV1.POST("leaderboard", ac.handleLeaderboard)
	gin.SetMode(gin.DebugMode)
	ac.engine = router
	return &ac, nil
}

func (ac *ApiController) Start(cnf config.Config, ctx context.Context) {
	var gracefulStop = make(chan os.Signal)
	srv := http.Server{
		Addr:    ":" + cnf.Api.Port,
		Handler: ac.engine,
	}
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		mlog.Printf("caught sig: %+v\n", sig)
		mlog.Printf("Server shutting down willingly")
		srv.Shutdown(ctx)
	}()
	if err := srv.ListenAndServe(); err != nil {
		ac.Shutdown(ctx)
	}
}

func (ac *ApiController) Shutdown(ctx context.Context) {
	log.Println("API shutting down")
	//TODO wait to complete requests
	ac.repo.ShutDown()
}
