package services


type BlacklistParameter struct{
	IPAddress string `json:"ip_address"`
	UserAgent string `json:"user_agent"`
}
