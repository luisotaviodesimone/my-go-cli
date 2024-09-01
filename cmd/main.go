package main

import (
	"github.com/luisotaviodesimone/my-go-cli/internal/cmds"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "lods [command]",
		Short: "Minha CLI para tarefas recorrentes",
	}
	rootCmd.AddCommand(cmds.Hello())
	rootCmd.AddCommand(cmds.FolderSizeCmd())
	rootCmd.AddCommand(cmds.GetCpfCmd())
	rootCmd.AddCommand(cmds.SetGitUserCmd())
	rootCmd.AddCommand(cmds.Speak())
	rootCmd.AddCommand(cmds.ApprovePrCmd())
	rootCmd.AddCommand(cmds.ClearNodeModulesCmd())
	rootCmd.AddCommand(cmds.Decode())
	rootCmd.AddCommand(cmds.Kube())
	rootCmd.AddCommand(cmds.Git())

	rootCmd.Execute()
}
