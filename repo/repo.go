package repo

import (
	"mage_test_case/config"
	"mage_test_case/mlog"
	"mage_test_case/model"
	"sort"
)

type Repo struct {
	mongo *Mongo
}

func New(cfg *config.Config) (*Repo, error) {

	mongo, err := NewMongo(cfg)
	if err != nil {
		mlog.Fatalf("Mongo init failed err : %+v", err)
		return nil, err
	}
	r := &Repo{
		mongo: mongo,
	}

	return r, nil
}

func (r *Repo) ShutDown() {
	mlog.Printf("Repo shutting down")
	// TODO wait connections
	r.mongo.ShutDown()
	mlog.Printf("Repo down.")
}

func (r *Repo) Register(req *model.RegisterRequest) (*model.RegisterResponse, error) {
	return r.mongo.registerMongo(req)
}

func (r *Repo) Login(req model.LoginRequest) (*model.LoginResponse, error) {
	return r.mongo.loginMongo(&req)
}

func (r *Repo) EndGame(req model.EndGameRequest) (*model.EndGameResponse, error) {
	players := req.Players
	sort.SliceStable(players, func(i, j int) bool {
		return players[i].Score < players[j].Score
	})
	for i := 0; i < len(players); i++ {
		players[i].Rank = i + 1
	}
	req.Players = players
	return r.mongo.endGameMongo(&req)
}

func (r *Repo) GetLeaderboard() (*model.LeaderboardResponse, error) {
	return r.mongo.getLeaderboardMongo()
}
