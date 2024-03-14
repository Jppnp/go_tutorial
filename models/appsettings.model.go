package models

type Configuration struct {
	DB_CONNECTION struct {
		CONNECTION_STRING string `json:"CONNECTION_STRING"`
		MAX_POOL          uint   `json:"MAX_POOL"`
		IDLE_POOL         uint   `json:"IDLE_POOL"`
		LIFE_TIME         uint   `json:"LIFE_TIME"`
	} `json:"DB_CONNECTION"`
	JWT_SECRET       string `json:"JWT_SECRET"`
	TOKEN_EXPIRATION int    `json:"TOKEN_EXPIRATION"`
}
