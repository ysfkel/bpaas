package login

import (
	"github.com/spf13/cobra"
	userSvc "github.com/ysfkel/bpaas/services/user"
)

type Login struct {
	Command *cobra.Command
}

func NewLogin() *Login {
	root := getRootCommand()

	login := &Login{
		Command: root,
	}

	return login
}

func (c *Login) GetLoginCommand() *cobra.Command {
	return c.Command
}

func getRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "login",
		Long:  `login`,
		Run: func(cmd *cobra.Command, args []string) {
			login(cmd, userSvc.NewUser())
		},
	}

	cmd.Flags().StringP("username", "u", "", "Username")
	cmd.Flags().StringP("password", "p", "", "Password")
	return cmd
}
