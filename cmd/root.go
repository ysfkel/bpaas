package cmd

import (
	"github.com/spf13/cobra"
	bpaas "github.com/ysfkel/bpaas/cmd/bpaas"
	"github.com/ysfkel/bpaas/cmd/bpaas/cluster"
	bpaasCluster "github.com/ysfkel/bpaas/cmd/bpaas/cluster"
	login "github.com/ysfkel/bpaas/cmd/bpaas/login"
	org "github.com/ysfkel/bpaas/cmd/bpaas/organization"
	orq "github.com/ysfkel/bpaas/cmd/bpaas/organization_request"
	usr "github.com/ysfkel/bpaas/cmd/bpaas/user"
)

func RunCli() {
	cluster := bpaasCluster.NewCluster()
	userCommand := usr.NewUser()
	loginCommand := login.NewLogin()
	orgCommand := org.NewOrganization()
	orqRequestCommand := orq.NewOrganizationRequest()
	bpaas := newBpaasRootCmd(cluster, userCommand, loginCommand, orgCommand, orqRequestCommand)
	bpaas.Execute()
}

func newBpaasRootCmd(cl cluster.ICluster,
	usr usr.IUser,
	lg login.ILogin,
	org org.IOrganization,
	orqRequest orq.IOrganizationRequest) *cobra.Command {

	bpaasCmd := bpaas.NewBpaascommand(cl, usr, lg, org, orqRequest)
	return bpaasCmd
}

func getBpaasCmdRoot() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "bpaas",
		Short: "bpaas",
		Long:  `bpaas`,
	}
	return cmd
}
