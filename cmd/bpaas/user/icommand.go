package user

import "github.com/spf13/cobra"

type IUser interface {
	GetUserCommand() *cobra.Command
}
