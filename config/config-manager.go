package config

import (
	"fmt"
	"sync"

	viperObj "github.com/spf13/viper"
)

var configManager *ConfigurationManager

var once sync.Once

type ConfigurationManager struct {
	configuration *Configuration
}

func NewConfigurationManager(configPath string, fileName string) (*ConfigurationManager, error) {

	var err error

	var config *Configuration

	once.Do(func() {
		config, err = createConfiguration(configPath, fileName)

		configManager = &ConfigurationManager{
			configuration: config,
		}
	})

	return configManager, err
}

func GetConfigManager() *ConfigurationManager {
	return configManager
}

func (cm *ConfigurationManager) GetApiVersion() string {
	return cm.configuration.ApiVersion
}

func (cm *ConfigurationManager) GetServer() string {
	return cm.configuration.Clusters[0].Cluster.Server
}

func (cm *ConfigurationManager) GetCertificateAuthority() string {
	return cm.configuration.Clusters[0].Cluster.CertificateAuthorityData
}

func (cm *ConfigurationManager) GetClientCertificate() string {
	return cm.configuration.Users[0].User.ClientCertificateData
}

func (cm *ConfigurationManager) GetClientKey() string {
	return cm.configuration.Users[0].User.ClientKeyData
}

func createConfiguration(configPath string, fileName string) (*Configuration, error) {

	var configuration Configuration

	viper, err := bindDefaultConfigurations(viperObj.New(), configPath, fileName)

	if err != nil {
		return nil, err
	}

	bindEnvironmentVariables(viper)

	err = viper.Unmarshal(&configuration)

	if err != nil {
		return nil, err
	}

	return &configuration, nil
}

func bindDefaultConfigurations(viper *viperObj.Viper, configPath string, fileName string) (*viperObj.Viper, error) {

	viper.SetConfigName(fileName)
	viper.AddConfigPath(configPath)
	viper.SetConfigType("yaml")

	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("..config error", err)
		return nil, err
	}

	return viper, nil
}

func bindEnvironmentVariables(viper *viperObj.Viper) {
	//viper.BindEnv("DB_HOST")
	// viper.BindEnv("DB_PORT")
	// viper.BindEnv("DB_NAME")
	// viper.BindEnv("DB_MOCK")
	// viper.BindEnv("DOCKER_REG")
	// viper.BindEnv("DB_USERNAME")
	// viper.BindEnv("DB_PASSWORD")
	// viper.BindEnv("AES_KEY")
	// viper.BindEnv("DB_PATH")
	// viper.BindEnv("DB_SSL_MODE")
	// viper.BindEnv("PLATFORM_ADMIN_EMAIL")
	// viper.BindEnv("PLATFORM_ADMIN_SECRET")
}
