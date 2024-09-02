package cmds

import (
	lodsgit "github.com/luisotaviodesimone/my-go-cli/internal/cmds/git"
	"github.com/spf13/cobra"
)

func Git() *cobra.Command {
	command := &cobra.Command{
		Use:   "go-git [command]",
		Short: "Uso do git através da minha cli",
	}
	command.AddCommand(lodsgit.CleanStaleBranches())

	return command
}
