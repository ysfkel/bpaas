package organization

import (
	"github.com/ysfkel/bpaas/config"
)

type Organization struct {
	Name      string //`json:"name"`
	ID        string //`json:"id"`
	CompanyID string //`json:"token"`
}

type OrganizationResponse struct {
	Data   []Organization `json:"data",yaml:"data"`
	Status string         `json:"status",yaml:"status"`
	Error  string         `json:"error",yaml:"error"`
}

type Response struct {
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
	Error  string      `json:"error"`
}

type Configuration struct {
	Token *config.Token
}
