package main

import (
	"github.com/luisotaviodesimone/my-go-cli/internal/cmds"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(cmds.Hello())
	rootCmd.AddCommand(cmds.FolderSizeCmd())
	rootCmd.AddCommand(cmds.GetCpfCmd())
	rootCmd.AddCommand(cmds.SetGitUserCmd())
	rootCmd.AddCommand(cmds.Speak())
	rootCmd.AddCommand((cmds.ApprovePrCmd()))

	rootCmd.Execute()
}
