package organization

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ysfkel/bpaas/services/organization"
)

func list(cmd *cobra.Command, org organization.IOrganization) error {

	cid, _ := cmd.Flags().GetString("company-id")

	if cid == "" {
		return printError("company id  is required. use -c ")
	}

	org.List(cid)

	return nil

}

func printError(msg string) error {
	fmtmsg := fmt.Sprintf("\n %s \n", msg)
	fmt.Println(fmtmsg)
	return nil
}
