package cmds

import (
	"github.com/luisotaviodesimone/my-go-cli/internal/cmds/kube"
	"github.com/spf13/cobra"
)

func Kube() *cobra.Command {
	command := &cobra.Command{
		Use:   "kube [command]",
		Short: "Gerenciamento do meu kubernetes e afins",
	}
	command.AddCommand(kube.ConfigKubeConfig())

	return command
}
