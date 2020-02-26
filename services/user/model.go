package user

import (
	"github.com/ysfkel/bpaas/config"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Token config.Token
}

type Response struct {
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
	Error  string      `json:"error"`
}

type Configuration struct {
	Token *config.Token
}
