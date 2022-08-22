package model

type Status string

const (
	StatusSuccess Status = "success"
	StatusFail    Status = "fail"
)

type PlayerScoreData struct {
	ID    int `json:"id" bson:"id"`
	Score int `json:"score" bson:"score"`
}

type PlayerRankData struct {
	ID   int `json:"id" bson:"id"`
	Rank int `json:"rank" bson:"rank"`
}
