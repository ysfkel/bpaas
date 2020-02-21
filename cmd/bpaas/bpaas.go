package bpaas

import (
	"github.com/spf13/cobra"
	"github.com/ysfkel/bpaas/cmd/bpaas/cluster"
)

func NewBpaascommand(cl cluster.ICluster) *cobra.Command {

	bpaasCmd := getBpaasCommand()
	clusterCommand := cl.GetClusterCommand()
	bpaasCmd.AddCommand(clusterCommand)

	return bpaasCmd
}

func getBpaasCommand() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "bpaas",
		Short: "bpaas",
		Long:  `bpaas`,
	}
	return cmd
}
