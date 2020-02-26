package organization_request

import "github.com/spf13/cobra"

type IOrganizationRequest interface {
	GetCreateCommand() *cobra.Command
	GetOrganizationRequestCommand() *cobra.Command
}
