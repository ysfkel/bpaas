package organization_request

import (
	"fmt"

	"github.com/spf13/cobra"
	orgReq "github.com/ysfkel/bpaas/services/organization_request"
)

func create(cmd *cobra.Command, org orgReq.IOrganizationRequest) error {

	dataPath, _ := cmd.Flags().GetString("data")

	if dataPath == "" {
		return printError("data file path  is required. use -d ")
	}

	err := org.Create(dataPath)

	if err != nil {
		fmt.Println(err)
	}

	return nil

}

func printError(msg string) error {
	fmtmsg := fmt.Sprintf("\n %s \n", msg)
	fmt.Println(fmtmsg)
	return nil
}
