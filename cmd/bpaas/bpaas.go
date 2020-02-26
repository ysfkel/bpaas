package bpaas

import (
	"github.com/spf13/cobra"
	"github.com/ysfkel/bpaas/cmd/bpaas/cluster"
	"github.com/ysfkel/bpaas/cmd/bpaas/login"
	"github.com/ysfkel/bpaas/cmd/bpaas/organization"
	orq "github.com/ysfkel/bpaas/cmd/bpaas/organization_request"
	"github.com/ysfkel/bpaas/cmd/bpaas/user"
)

func NewBpaascommand(cl cluster.ICluster, usr user.IUser, lg login.ILogin, org organization.IOrganization, orq orq.IOrganizationRequest) *cobra.Command {

	bpaasCmd := getBpaasCommand()
	clusterCommand := cl.GetClusterCommand()
	bpaasCmd.AddCommand(clusterCommand)
	//user
	// userCommand := usrGetUserCommand()
	// bpaasCmd.AddCommand(userCommand)
	//login
	loginCommand := lg.GetLoginCommand()
	bpaasCmd.AddCommand(loginCommand)
	//org
	orgCommand := org.GetOrganizationCommand()
	bpaasCmd.AddCommand(orgCommand)
	//orq request
	orgReqCommand := orq.GetOrganizationRequestCommand()
	bpaasCmd.AddCommand(orgReqCommand)
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
