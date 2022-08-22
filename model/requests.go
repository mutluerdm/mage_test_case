package model

type RegisterRequest struct {
	Id       int64  `json:"id" bson:"id"`
	Username string `json:"username" bson:"username" binding:"required,alphanum,min=3,max=10"`
	Password string `json:"password" bson:"password" binding:"required,alphanum,min=3,max=10"`
	Score    int    `json:"score" bson:"score"`
}

type LoginRequest struct {
	Username string `json:"username" bson:"username" binding:"required,alphanum,min=3,max=10"`
	Password string `json:"password" bson:"password" binding:"required,alphanum,min=3,max=10"`
}

type EndGameRequest struct {
	Players []PlayerEndGameData `json:"players" bson:"players" binding:"required"`
}
type PlayerEndGameData struct {
	ID    int `json:"id" bson:"id"`
	Score int `json:"score" bson:"score"`
	Rank  int `json:"rank"  bson:"rank"`
}

func (endgame *EndGameRequest) ToPlayerRankData() []PlayerRankData {
	res := []PlayerRankData{}
	for _, v := range endgame.Players {
		res = append(res, PlayerRankData{ID: v.ID, Rank: v.Rank})
	}
	return res
}

func (endgame *EndGameRequest) ToPlayerScoreData() []PlayerScoreData {
	res := []PlayerScoreData{}
	for _, v := range endgame.Players {
		res = append(res, PlayerScoreData{ID: v.ID, Score: v.Score})
	}
	return res
}
