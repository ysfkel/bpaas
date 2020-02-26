package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ysfkel/bpaas/config"
	"github.com/ysfkel/bpaas/filesystem"
	"github.com/ysfkel/bpaas/http"
)

func NewUser() *User {
	return &User{}
}

func (u *User) Login(username string, password string) error {

	const url = "http://localhost:5000/v1/api/users/auth"

	user := &User{Email: username, Password: password}

	userJson, err := json.Marshal(user)

	if err != nil {
		return err
	}

	userBuffer := bytes.NewBuffer(userJson)

	httpRequest, err := http.NewHttpRequest("POST", url, "", userBuffer)

	if err != nil {
		fmt.Println(err)
		return err
	}

	response, err := httpRequest.Post()

	if err != nil {
		fmt.Println(err)
		return err
	}

	if response.Data == nil {
		return errors.New("no login token returned")
	}

	token, ok := response.Data.(string)

	if !ok {
		return errors.New("error converting login token")
	}

	userConfig := Configuration{
		Token: &config.Token{
			Jwt: token,
		},
	}
	yamlData, err := filesystem.CreateYAML(&userConfig)

	if err != nil {
		fmt.Println(err)
		return err
	}

	configPath, err := filesystem.GetUserConfigPath()

	if err != nil {
		fmt.Println(err)
		return err
	}

	err = filesystem.WriteToYAMLFile(yamlData, "auth", configPath)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}
