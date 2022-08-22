package model

type SuccessResponse struct {
	Status    Status      `json:"status"`
	Timestamp int64       `json:"timestamp"`
	Result    interface{} `json:"result"`
}

type FailResponse struct {
	Status Status `json:"status"`
	Error  string `json:"Error"`
}

type RegisterResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type EndGameResponse []PlayerScoreData

type LeaderboardResponse []PlayerRankData
