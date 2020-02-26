package model

import "github.com/ysfkel/bpaas/config"

type OrganizationRequestCreate struct {
	CompanyID             string `json:"company_id"`
	CompanyName           string `json:"company_name"`
	OrgName               string `json:"org_name"`
	Description           string `json:"description"`
	ChaincodeCount        uint32 `json:"chaincode_count"`
	PeerCount             uint32 `json:"peer_count"`
	OrdererCount          uint32 `json:"orderer_count"`
	ConsortiumName        string `json:"consortium_name"`
	DbBackend             string `json:"db_backend"`
	ConsortiumDescription string `json:"consortium_description"`
	UserName              string `json:"user_name"`
	UserEmail             string `json:"user_email"`
}

type Configuration struct {
	Token *config.Token
}
