package repo

import (
	"context"
	"mage_test_case/config"
	"mage_test_case/mlog"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repo struct {
	mongo *mongo.Client
}

func New(cfg *config.Config) (error, *Repo) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//TODO configure connections
	mClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		mlog.PrintErrf("Mongo Connection Error : %+v", err)
		return err, nil
	}
	r := &Repo{mongo: mClient}
	//r.CheckVersions()

	return nil, r
}

func (r *Repo) ShutDown() {
	mlog.Printf("Repo shutting down")
	// TODO wait connections and
	mlog.Printf("Repo down.")
}
