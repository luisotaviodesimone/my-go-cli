package cmds

import (
	"github.com/luisotaviodesimone/my-go-cli/internal/cmds/typerace"
	"github.com/spf13/cobra"
)

func Typerace() *cobra.Command {
	command := &cobra.Command{
		Use:   "typerace [command]",
		Short: "Cópia fajuta do jogo typeracer",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
      objective := typerace.GetQuote()
			typerace.StartTyperace(objective.Text)
		},
	}


	return command
}
