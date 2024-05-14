package config

type TokenType uint

const (
	ACCESS  TokenType = 0
	REFRESH TokenType = 1
)

type Status uint

const (
	IN_PROGRESS 		Status = 0
	TEMPORARILY_STOPPED Status = 1
	COMPLETED 			Status = 2
)

