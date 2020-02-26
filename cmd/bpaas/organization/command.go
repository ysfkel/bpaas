package organization

import (
	"github.com/spf13/cobra"
	"github.com/ysfkel/bpaas/services/organization"
)

type Organization struct {
	Command *cobra.Command
}

func NewOrganization() *Organization {
	root := GetRoot()
	// root.AddCommand()
	org := &Organization{
		Command: root,
	}

	//add create command
	org.Command.AddCommand(org.GetListCommand())

	return org
}

func GetRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "org",
		Short: "org",
		Long:  `org`,
	}
	return cmd
}

func (c *Organization) GetOrganizationCommand() *cobra.Command {
	return c.Command
}

func (c *Organization) GetListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list",
		Long:  `list`,
		Run: func(cmd *cobra.Command, args []string) {
			list(cmd, organization.NewOrganization())
		},
	}

	cmd.Flags().StringP("company-id", "c", "", "company-id")
	return cmd
}
