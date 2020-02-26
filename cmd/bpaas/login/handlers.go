package login

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ysfkel/bpaas/services/user"
)

func login(cmd *cobra.Command, u user.IUser) error {

	username, _ := cmd.Flags().GetString("username")

	if username == "" {
		return printError("username is required. use -u ")
	}

	password, _ := cmd.Flags().GetString("password")

	if password == "" {
		return printError("password is required. use -u ")
	}

	fmt.Println("hello user")

	u.Login(username, password)

	return nil

}

func printError(msg string) error {
	fmtmsg := fmt.Sprintf("\n %s \n", msg)
	fmt.Println(fmtmsg)
	return nil
}
