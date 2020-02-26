package user

import (
	"github.com/spf13/cobra"
)

type User struct {
	Command *cobra.Command
}

func NewUser() *User {
	root := GetRoot()
	// root.AddCommand()
	user := &User{
		Command: root,
	}

	//add create command
	//user.Command.AddCommand(user.LoginCommand())

	return user
}

func GetRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "user",
		Long:  `user`,
	}
	return cmd
}

func (c *User) GetUserCommand() *cobra.Command {
	return c.Command
}
