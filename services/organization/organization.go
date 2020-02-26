package organization

import (
	"encoding/json"
	"fmt"

	"github.com/ysfkel/bpaas/config"
	"github.com/ysfkel/bpaas/filesystem"
	"github.com/ysfkel/bpaas/http"
)

type Token struct {
	Token config.Token
}

func NewOrganization() *Organization {
	return &Organization{}
}

func (u *Organization) List(companyID string) error {

	var url = "http://localhost:5000/v1/api/organizations?company_id=" + companyID

	//"http://localhost:5000/v1/api/organizations?company_id=c1e18f33-32a8-455c-ab1c-d805f1a4fcc0"

	path, err := filesystem.GetUserConfigPath()

	if err != nil {
		return err
	}

	configManager, err := config.NewConfigurationManager(path, "auth")

	if err != nil {
		return err
	}

	token := configManager.GetJWT()

	httpRequest, err := http.NewHttpRequest("GET", url, token, nil)

	if err != nil {
		fmt.Println(err)
		return err
	}

	response, err := httpRequest.Execute()

	if err != nil {
		fmt.Println(err)
		return err
	}

	var rs OrganizationResponse

	json.Unmarshal(response, &rs)

	PrintOut(rs.Data)

	return nil

}

// func (u *Organization) Create(username string, password string) error {

// 	const url = "http://localhost:5000/v1/api/users/auth"

// 	user := &Organization{Email: username, Password: password}

// 	userJson, err := json.Marshal(user)

// 	if err != nil {
// 		return err
// 	}

// 	userBuffer := bytes.NewBuffer(userJson)

// 	httpRequest, err := http.NewHttpRequest("POST", url, "", userBuffer)

// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	response, err := httpRequest.Post()

// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	if response.Data == nil {
// 		return errors.New("no login token returned")
// 	}

// 	token, ok := response.Data.(string)

// 	if !ok {
// 		return errors.New("error converting login token")
// 	}

// 	userConfig := Configuration{
// 		Token: &config.Token{
// 			Jwt: token,
// 		},
// 	}
// 	yamlData, err := filesystem.CreateYAML(&userConfig)

// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	configPath, err := filesystem.GetOrganizationConfigPath()

// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	err = filesystem.WriteToYAMLFile(yamlData, "auth", configPath)

// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	return nil

// }
