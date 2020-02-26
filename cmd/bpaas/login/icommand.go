package login

import "github.com/spf13/cobra"

type ILogin interface {
	GetLoginCommand() *cobra.Command
}
