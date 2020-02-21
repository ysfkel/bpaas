package filesystem

import (
	"fmt"
	"os"
)

const (
	configDirPath = "/.bpaas/config/clusters"
)

func CreateConfigDirIfNotExists() error {

	configPath, err := GetConfigPath()

	if err != nil {
		fmt.Println("...error retrieving home path", err)
		return err
	}

	if IsNotExistsConfigDirectory(configPath) {
		fmt.Println("...creating  ", configPath)
		return CreateDirectory(configPath)
	}

	if err != nil && !os.IsNotExist(err) {
		fmt.Println("...error ", err)
		return err
	} else if err != nil {
		fmt.Println("...error  ", err)
		return err
	}

	return nil
}

func IsNotExistsConfigDirectory(configPath string) bool {

	_, err := os.Stat(configPath)

	if err != nil && os.IsNotExist(err) {
		return true
	}

	return false
}

func GetConfigPath() (string, error) {

	home, err := GetUserHomeDirectory()

	if err != nil {
		fmt.Println("...error retrieving home path", err)
		return "", err
	}

	path := fmt.Sprintf("%s/%s", home, configDirPath)

	return path, nil
}
func GetUserHomeDirectory() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return homeDir, nil
}

func CreateDirectory(path string) error {

	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		return err
	}

	return nil

}
