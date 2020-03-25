package cluster

import (
	"fmt"

	"github.com/spf13/cobra"
	k8s "github.com/ysfkel/kubernetes/provider"
)

func createCluster(cmd *cobra.Command, KubernetesService k8s.IKubernetesService) error {

	name, _ := cmd.Flags().GetString("name")

	if name == "" {
		return printError("cluster name is required. use -n ")
	}

	err := KubernetesService.CreateCluster(name)

	if err != nil {
		printError("could not create cluster")
		return nil
	}

	return nil
}

func destroyCluster(cmd *cobra.Command, KubernetesService k8s.IKubernetesService) error {

	name, _ := cmd.Flags().GetString("name")
	if name == "" {
		return printError("cluster name is equired. use -n ")
	}

	err := KubernetesService.DestroyCluster(name)

	if err != nil {
		printError(fmt.Sprintf("could not delete cluster %s", err))
		return nil
	}
	return nil
}

func getCertificateAuthority(cmd *cobra.Command, KubernetesService k8s.IKubernetesService) error {

	name, _ := cmd.Flags().GetString("name")

	if name == "" {
		return printError("cluster name is equired. use -n ")
	}

	ca, err := KubernetesService.GetCertificateAuthority(name)

	if err != nil {
		return printError(fmt.Sprintf("could not fetch certifate authority for cluster %s %s", name, err))
	}

	fmt.Println(ca)

	return nil
}

func getClientCertificate(cmd *cobra.Command, KubernetesService k8s.IKubernetesService) error {

	name, _ := cmd.Flags().GetString("name")

	if name == "" {
		return printError("cluster name is equired. use -n ")
	}

	ca, err := KubernetesService.GetClientCertificate(name)

	if err != nil {
		return printError(fmt.Sprintf("could not fetch client certifate for cluster %s %s", name, err))
	}

	fmt.Println(ca)

	return nil
}

func getClientKey(cmd *cobra.Command, KubernetesService k8s.IKubernetesService) error {

	name, _ := cmd.Flags().GetString("name")

	if name == "" {
		return printError("cluster name is equired. use -n ")
	}

	ck, err := KubernetesService.GetClientKey(name)

	if err != nil {
		return printError(fmt.Sprintf("could not fetch client key for cluster %s %s", name, err))
	}

	fmt.Println(ck)

	return nil
}

func getClusterIP(cmd *cobra.Command, KubernetesService k8s.IKubernetesService) error {

	name, _ := cmd.Flags().GetString("name")

	if name == "" {
		return printError("cluster name is equired, use -n")
	}

	ck, err := KubernetesService.GetClusterIP(name)

	if err != nil {
		return printError(fmt.Sprintf("could not fetch cluster ip  for cluster %s %s", name, err))
	}

	fmt.Println(ck)

	return nil
}

func printError(msg string) error {
	fmtmsg := fmt.Sprintf("\n %s \n", msg)
	fmt.Println(fmtmsg)
	return nil
}
