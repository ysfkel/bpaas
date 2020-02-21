package cluster

import "github.com/spf13/cobra"

type ICluster interface {
	GetClusterCommand() *cobra.Command
	GetCreateCommand() *cobra.Command
	GetDeleteCommand() *cobra.Command
	GetCertificateAuthorityCommand() *cobra.Command
	GetClientCertificateCommand() *cobra.Command
	GetClientKeyCommand() *cobra.Command
	GetClusterIPCommand() *cobra.Command
}
