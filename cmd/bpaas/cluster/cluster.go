package cluster

import (
	"fmt"

	"github.com/spf13/cobra"
	k8sKind "github.com/ysfkel/bpaas/kubernetes/kind"
)

type Cluster struct {
	Command *cobra.Command
}

func NewCluster() *Cluster {
	root := ClusterRoot()
	// root.AddCommand()
	cluster := &Cluster{
		Command: root,
	}

	//add create command
	cluster.Command.AddCommand(cluster.GetCreateCommand())
	//add destroy command
	cluster.Command.AddCommand(cluster.GetDeleteCommand())

	getCommand := cluster.GetCommand()
	getCommand.AddCommand(cluster.GetCertificateAuthorityCommand())
	getCommand.AddCommand(cluster.GetClientCertificateCommand())
	getCommand.AddCommand(cluster.GetClientKeyCommand())
	getCommand.AddCommand(cluster.GetClusterIPCommand())
	//add get command
	cluster.Command.AddCommand(getCommand)

	return cluster
}

func ClusterRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cluster",
		Short: "cluster",
		Long:  `cluster`,
	}
	return cmd
}

func (c *Cluster) GetClusterCommand() *cobra.Command {
	return c.Command
}

func (c *Cluster) GetCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create",
		Long:  `create`,
		Run: func(cmd *cobra.Command, args []string) {
			createCluster(cmd, k8sKind.NewKubernetesService())
		},
	}

	cmd.Flags().StringP("name", "n", "", "Name of the folder you want to create.")

	return cmd
}

func (c *Cluster) GetDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete",
		Long:  `delete`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("..destroying cluster")
			destroyCluster(cmd, k8sKind.NewKubernetesService())
		},
	}
	cmd.Flags().StringP("name", "n", "", "Name of the folder you want to create.")
	return cmd
}

func (c *Cluster) GetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get",
		Long:  `get`,
	}
	cmd.Flags().StringP("name", "n", "", "Name of the folder you want to create.")
	return cmd
}

func (c *Cluster) GetCertificateAuthorityCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ca",
		Short: "certificate-authority",
		Long:  `certificate-authority`,
		Run: func(cmd *cobra.Command, args []string) {
			getCertificateAuthority(cmd, k8sKind.NewKubernetesService())
		},
	}
	cmd.Flags().StringP("name", "n", "", "Name of the folder you want to create.")
	return cmd
}
func (c *Cluster) GetClientCertificateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cc",
		Short: "client-certificate",
		Long:  `client-certificate`,
		Run: func(cmd *cobra.Command, args []string) {
			getClientCertificate(cmd, k8sKind.NewKubernetesService())
		},
	}
	cmd.Flags().StringP("name", "n", "", "Name of the folder you want to create.")
	return cmd
}

func (c *Cluster) GetClientKeyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ck",
		Short: "client-key",
		Long:  `client-key`,
		Run: func(cmd *cobra.Command, args []string) {
			getClientKey(cmd, k8sKind.NewKubernetesService())
		},
	}
	cmd.Flags().StringP("name", "n", "", "Name of the folder you want to create.")
	return cmd
}

func (c *Cluster) GetClusterIPCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ip",
		Short: "Retrieves the cluster IP",
		Long:  "Retrieves the cluster IP",
		Run: func(cmd *cobra.Command, args []string) {
			getClusterIP(cmd, k8sKind.NewKubernetesService())
		},
	}
	cmd.Flags().StringP("name", "n", "", "Name of the folder you want to create.")
	return cmd
}
