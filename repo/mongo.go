package repo

import (
	"context"
	"errors"
	"log"
	"mage_test_case/config"
	"mage_test_case/mlog"
	"mage_test_case/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	colCounters    = "counters"
	colUsers       = "users"
	colUserScores  = "user_scores"
	colGameLog     = "game_log"
	colLeaderboard = "leaderboard"
)

type Mongo struct {
	mongoClient   *mongo.Client
	mongoDatabase *mongo.Database
	autoInc       *AutoInc
}

func NewMongo(cfg *config.Config) (*Mongo, error) {
	clientOptions := options.Client().
		ApplyURI(cfg.Mongo.Url)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	mainDB := client.Database(cfg.Mongo.Database)
	usernameUniqueIndex := mongo.IndexModel{
		Keys: bson.M{
			"username": 1, // index in ascending order
		}, Options: options.Index().SetUnique(true),
	}
	_, err = mainDB.Collection(colUsers).Indexes().CreateOne(context.Background(), usernameUniqueIndex)
	if err != nil {
		mlog.PrintErrf("Create user unique Index Err : %+v", err)
	}
	leaderboardIndex := mongo.IndexModel{
		Keys: bson.M{
			"rank": -1, // index in desc order
		}, Options: options.Index(),
	}
	_, err = mainDB.Collection(colLeaderboard).Indexes().CreateOne(context.Background(), leaderboardIndex)
	if err != nil {
		mlog.PrintErrf("Create leaderobard Index Err : %+v", err)
	}
	userScoresIndex := mongo.IndexModel{
		Keys: bson.M{
			"score": -1, // index in desc order
		}, Options: options.Index(),
	}
	_, err = mainDB.Collection(colUserScores).Indexes().CreateOne(context.Background(), userScoresIndex)
	if err != nil {
		mlog.PrintErrf("Create user scores Index Err : %+v", err)
	}
	mongo := Mongo{
		mongoClient:   client,
		mongoDatabase: mainDB,
		autoInc:       NewUtoInc(mainDB.Collection(colCounters))}
	return &mongo, nil
}

func (m *Mongo) ShutDown() {
	//TODO wait tasks
}

func (m *Mongo) registerMongo(req *model.RegisterRequest) (*model.RegisterResponse, error) {
	req.Id = (int64)(m.autoInc.Next(colUsers))
	_, err := m.mongoDatabase.Collection(colUsers).InsertOne(context.TODO(), &req)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, errors.New("username must be unique")
		}
		return nil, err
	}
	_, err = m.mongoDatabase.Collection(colUserScores).InsertOne(context.TODO(), &model.PlayerScoreData{ID: int(req.Id)})
	if err != nil {
		return nil, err
	}
	_, err = m.mongoDatabase.Collection(colLeaderboard).InsertOne(context.TODO(), &model.PlayerRankData{ID: int(req.Id)})
	if err != nil {
		return nil, err
	}
	return &model.RegisterResponse{ID: req.Id, Username: req.Username, Password: req.Password}, nil
}

func (m *Mongo) loginMongo(req *model.LoginRequest) (*model.LoginResponse, error) {
	//opts := options.FindOne().SetProjection(bson.M{"id": 1, "username": 1})
	res := m.mongoDatabase.Collection(colUsers).FindOne(context.Background(), req)
	if res.Err() != nil {
		mlog.PrintErrf("Login Error err : %+v", res.Err())
		if res.Err() == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, res.Err()
	}
	response := &model.LoginResponse{}
	res.Decode(response)
	mlog.Printf("login result : %+v", response)
	return response, nil
}

func (m *Mongo) getLeaderboardMongo() (*model.LeaderboardResponse, error) {
	opts := options.Find().SetSort(bson.M{"rank": -1})
	cur, err := m.mongoDatabase.Collection(colLeaderboard).Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		mlog.PrintErrf("mongo getleaderborad err : %+v", err)
		return nil, err
	}
	result := model.LeaderboardResponse{}
	for cur.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem model.PlayerRankData
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, elem)
	}
	return &result, err
}

func (m *Mongo) endGameMongo(req *model.EndGameRequest) (*model.EndGameResponse, error) {
	_, err := m.mongoDatabase.Collection(colGameLog).InsertOne(context.TODO(), req)
	if err != nil {
		mlog.PrintErrf("endgame mongo insert game_log err : %+v", err)
		return nil, err
	}
	res := m.updateScoresMongo(req.ToPlayerScoreData())
	m.updateRanksMongo(req.ToPlayerRankData())
	return (*model.EndGameResponse)(&res), nil
}

func (m *Mongo) updateScoresMongo(req []model.PlayerScoreData) []model.PlayerScoreData {
	retunDoc := options.After
	upsert := false
	opts := &options.FindOneAndUpdateOptions{ReturnDocument: &retunDoc, Upsert: &upsert}
	playersUpdatedScores := []model.PlayerScoreData{}
	for _, v := range req {
		filter := bson.M{"id": v.ID}
		res := m.mongoDatabase.Collection(colUserScores).FindOneAndUpdate(context.Background(), filter, bson.M{"$inc": bson.M{"score": v.Score}}, opts)
		if res.Err() != nil {
			mlog.Println("mongo update scores error:", res.Err())
		} else {
			newScoreData := model.PlayerScoreData{}
			_ = res.Decode(&newScoreData)
			playersUpdatedScores = append(playersUpdatedScores, newScoreData)
		}

	}
	mlog.Printf("Updated Scores : %+v", playersUpdatedScores)
	return playersUpdatedScores
}

func (m *Mongo) updateRanksMongo(req []model.PlayerRankData) error {
	retunDoc := options.After
	upsert := false
	opts := &options.FindOneAndUpdateOptions{ReturnDocument: &retunDoc, Upsert: &upsert}
	for _, v := range req {
		filter := bson.M{"id": v.ID}
		res := m.mongoDatabase.Collection(colLeaderboard).FindOneAndUpdate(context.Background(), filter, bson.M{"$inc": bson.M{"rank": v.Rank}}, opts)
		if res.Err() != nil {
			mlog.Println("mongo update rank error:", res.Err())
		}
	}
	return nil
}
