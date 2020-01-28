package handler

type Response struct {
	ResponseTime string `json:"response_time"`
	IsBlacklist bool `json:"is_blacklist"`
}
