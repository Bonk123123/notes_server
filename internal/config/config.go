package config

type TokenType string

const (
	ACCESS  TokenType = "ACCESS_TOKEN_TTL"
	REFRESH TokenType = "REFRESH_TOKEN_TTL"
)

type Status uint

const (
	IN_PROGRESS 		Status = 0
	TEMPORARILY_STOPPED Status = 1
	COMPLETED 			Status = 2
)

