package cmd

import (
	"github.com/spf13/cobra"
	bpaas "github.com/ysfkel/bpaas/cmd/bpaas"
	"github.com/ysfkel/bpaas/cmd/bpaas/cluster"
	bpaasCluster "github.com/ysfkel/bpaas/cmd/bpaas/cluster"
)

func RunCli() {
	cluster := bpaasCluster.NewCluster()
	bpaas := newBpaasRootCmd(cluster)
	bpaas.Execute()
}

func newBpaasRootCmd(cl cluster.ICluster) *cobra.Command {

	//bpaasCmdRoot := getBpaasCmdRoot()
	bpaasCmd := bpaas.NewBpaascommand(cl)
	//	bpaasCmdRoot.AddCommand(bpaasCmd)
	return bpaasCmd //bpaasCmdRoot
}

func getBpaasCmdRoot() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "bpaas-",
		Short: "bpaas-",
		Long:  `bpaas-`,
	}
	return cmd
}
