package cmds

import (
	"fmt"

	"github.com/luisotaviodesimone/my-go-cli/internal/colors"
	"github.com/spf13/cobra"
)

func Hello() *cobra.Command {
	return &cobra.Command{
		Use:   "hello [name]",
		Short: "retorna Olá + name passado",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Olá %s%s%s\n", colors.Purple, args[0], colors.Reset)
		},
	}
}
