package organization

import "github.com/spf13/cobra"

type IOrganization interface {
	GetListCommand() *cobra.Command
	GetOrganizationCommand() *cobra.Command
}
