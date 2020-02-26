package organization_request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/ysfkel/bpaas/config"
	"github.com/ysfkel/bpaas/filesystem"
	"github.com/ysfkel/bpaas/http"
)

type OrganizationRequest struct {
}

func NewOrganizationRequest() *OrganizationRequest {
	return &OrganizationRequest{}
}

func getJWT() (string, error) {

	path, err := filesystem.GetUserConfigPath()

	if err != nil {
		return "", err
	}

	configManager, err := config.NewConfigurationManager(path, "auth")

	if err != nil {
		return "", err
	}

	token := configManager.GetJWT()

	return token, nil
}

func (u *OrganizationRequest) List(companyID string) error {

	//var url = "http://localhost:5000/v1/api/organizations?company_id=" + companyID

	// path, err := filesystem.GetUserConfigPath()

	// if err != nil {
	// 	return err
	// }

	// configManager, err := config.NewConfigurationManager(path, "auth")

	// if err != nil {
	// 	return err
	// }

	// token := configManager.GetJWT()

	// httpRequest, err := http.NewHttpRequest("GET", url, token, nil)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	// response, err := httpRequest.Execute()

	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	// var rs OrganizationResponse

	// json.Unmarshal(response, &rs)

	//PrintOut(rs.Data)

	return nil

}

func (u *OrganizationRequest) Create(dataPath string) error {

	const url = "http://localhost:5000/v1/api/organization-requests"

	filename := filepath.Base(dataPath)

	path := filepath.Dir(dataPath)

	cm, err := config.NewConfigurationManager(path, filename)

	orgRequest := cm.GetOrganizationRequest()

	orgJson, err := json.Marshal(orgRequest)

	if err != nil {
		return err
	}

	orgBuffer := bytes.NewBuffer(orgJson)

	token, err := getJWT()

	if err != nil {
		return err
	}

	httpRequest, err := http.NewHttpRequest("POST", url, token, orgBuffer)

	if err != nil {
		fmt.Println(err)
		return err
	}

	response, err := httpRequest.Execute()

	if err != nil {
		fmt.Println(err)
		return err
	}

	var rs http.Response

	json.Unmarshal(response, &rs)

	if rs.Error != "" {
		fmt.Println(rs.Error)
	}

	return nil

}
