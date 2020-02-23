package kind

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ysfkel/bpaas/config"
	"github.com/ysfkel/bpaas/filesystem"
	k8s "github.com/ysfkel/bpaas/kubernetes"
)

type KubernetesService struct {
}

func NewKubernetesService() k8s.IKubernetesService {

	return &KubernetesService{}
}

func (k *KubernetesService) CreateCluster(name string) error {

	err := filesystem.CreateConfigDirIfNotExists()

	if err != nil {
		return err
	}

	e := exec.Command("kind", "create", "cluster", "--name", name).Run()

	if e != nil {
		return e
	} else {
		return writeConfig(name)
	}
	return nil
}

func (k *KubernetesService) DestroyCluster(clusterName string) error {

	e := exec.Command("kind", "delete", "cluster", "--name", clusterName).Run()

	if e != nil {
		return e
	} else {
		return removeFile(clusterName)
	}
}

func (k *KubernetesService) GetCertificateAuthority(clusterName string) (string, error) {

	configDir, err := filesystem.GetConfigPath()

	if err != nil {
		return "", err
	}

	c, err := config.NewConfigurationManager(configDir, clusterName)

	if err != nil {
		return "", err
	}

	return c.GetCertificateAuthority(), nil

}
func (k *KubernetesService) GetClientCertificate(clusterName string) (string, error) {

	configDir, err := filesystem.GetConfigPath()

	if err != nil {
		return "", err
	}

	c, err := config.NewConfigurationManager(configDir, clusterName)

	if err != nil {
		return "", err
	}

	return c.GetClientCertificate(), nil
}
func (k *KubernetesService) GetClientKey(clusterName string) (string, error) {

	configDir, err := filesystem.GetConfigPath()

	if err != nil {
		return "", err
	}

	c, err := config.NewConfigurationManager(configDir, clusterName)

	if err != nil {
		return "", err
	}

	return c.GetClientKey(), nil
}

func (k *KubernetesService) GetClusterIP(clusterName string) (string, error) {

	configDir, err := filesystem.GetConfigPath()

	if err != nil {
		return "", err
	}

	c, err := config.NewConfigurationManager(configDir, clusterName)

	if err != nil {
		return "", err
	}

	fmt.Println("---init conf ", c, err)

	return c.GetServer(), nil
}

func writeConfig(clusterName string) error {

	result, err := exec.Command("kind", "get", "kubeconfig", "--name", clusterName).CombinedOutput()

	if err != nil {
		return err
	}
	return writeToYAMLFile(result, clusterName)
}

func writeToYAMLFile(data []byte, fileName string) error {

	configDir, err := filesystem.GetConfigPath()

	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("%v/%v.yaml", configDir, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	fmt.Println("..writing config to ", filePath)
	_, err = file.WriteString(string(data))

	if err != nil {
		file.Close()
		return err
	}
	return nil
}

func removeFile(fileName string) error {

	configDir, err := filesystem.GetConfigPath()

	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("%v/%v.yaml", configDir, fileName)

	err = os.Remove(filePath)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	return nil
}
