package cluster

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	k8s "github.com/ysfkel/bpaas/kubernetes"
)

func createCluster(cmd *cobra.Command, KubernetesService k8s.IKubernetesService) error {

	name, _ := cmd.Flags().GetString("name")

	if name == "" {
		fmt.Println("cluster name is equired. use -n ")
		return errors.New("cluster name is equired. use -n ")
	}

	err := KubernetesService.CreateCluster(name)

	if err != nil {
		fmt.Printf("could not create cluster")
		return err
	}

	return nil
}

func destroyCluster(cmd *cobra.Command, KubernetesService k8s.IKubernetesService) error {

	name, _ := cmd.Flags().GetString("name")
	if name == "" {
		fmt.Println("cluster name is equired. use -n ")
		return errors.New("cluster name is equired. use -n ")
	}

	err := KubernetesService.DestroyCluster(name)

	if err != nil {
		fmt.Printf("could not create cluster")
		return err
	}
	return nil
}

func getCertificateAuthority(cmd *cobra.Command, KubernetesService k8s.IKubernetesService) error {

	name, _ := cmd.Flags().GetString("name")

	if name == "" {
		fmt.Println("cluster name is equired. use -n ")
		return errors.New("cluster name is equired. use -n ")
	}

	ca, err := KubernetesService.GetCertificateAuthority(name)

	if err != nil {
		fmt.Printf("could not fetch certifate authgority for cluster %s", name)
		return fmt.Errorf("could not fetch certifate authgority for cluster %s", name)
	}

	fmt.Println(ca)

	return nil
}

func getClientCertificate(cmd *cobra.Command, KubernetesService k8s.IKubernetesService) error {

	name, _ := cmd.Flags().GetString("name")

	if name == "" {
		fmt.Println("cluster name is equired. use -n ")
		return errors.New("cluster name is equired. use -n ")
	}

	ca, err := KubernetesService.GetClientCertificate(name)

	if err != nil {
		fmt.Printf("could not fetch client certifate for cluster %s", name)
		return fmt.Errorf("could not fetch client certifate for cluster %s", name)
	}

	fmt.Println(ca)

	return nil
}

func getClientKey(cmd *cobra.Command, KubernetesService k8s.IKubernetesService) error {

	name, _ := cmd.Flags().GetString("name")

	if name == "" {
		fmt.Println("cluster name is equired. use -n ")
		return errors.New("cluster name is equired. use -n ")
	}

	ck, err := KubernetesService.GetClientKey(name)

	if err != nil {
		fmt.Printf("could not fetch client key for cluster %s", name)
		return fmt.Errorf("could not fetch client key for cluster %s", name)
	}

	fmt.Println(ck)

	return nil
}

func getClusterIP(cmd *cobra.Command, KubernetesService k8s.IKubernetesService) error {

	name, _ := cmd.Flags().GetString("name")

	if name == "" {
		fmt.Println("cluster name is equired. use -n ")
		return errors.New("cluster name is equired. use -n ")
	}

	ck, err := KubernetesService.GetClusterIP(name)

	if err != nil {
		fmt.Printf("could not fetch cluster ip  for cluster %s", name)
		return fmt.Errorf("could not fetch luster ip  for cluster %s", name)
	}

	fmt.Println(ck)

	return nil
}
