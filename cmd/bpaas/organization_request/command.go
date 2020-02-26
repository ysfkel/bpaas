package organization_request

import (
	"github.com/spf13/cobra"
	orq "github.com/ysfkel/bpaas/services/organization_request"
)

type OrganizationRequest struct {
	Command *cobra.Command
}

func NewOrganizationRequest() *OrganizationRequest {
	root := GetRoot()
	// root.AddCommand()
	orgReq := &OrganizationRequest{
		Command: root,
	}

	//add create command
	orgReq.Command.AddCommand(orgReq.GetListCommand())
	orgReq.Command.AddCommand(orgReq.GetCreateCommand())
	return orgReq
}

func GetRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "org-request",
		Short: "org-request",
		Long:  `org-request`,
	}
	return cmd
}

func (c *OrganizationRequest) GetOrganizationRequestCommand() *cobra.Command {
	return c.Command
}

func (c *OrganizationRequest) GetCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create",
		Long:  `create`,
		Run: func(cmd *cobra.Command, args []string) {
			create(cmd, orq.NewOrganizationRequest())
		},
	}

	cmd.Flags().StringP("data", "d", "", "data")
	return cmd
}

func (c *OrganizationRequest) GetListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list",
		Long:  `list`,
		Run: func(cmd *cobra.Command, args []string) {
			//list(cmd, organization.NewOrganization())
		},
	}

	cmd.Flags().StringP("company-id", "c", "", "company id")
	return cmd
}
