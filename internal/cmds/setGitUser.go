package cmds

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

func setGitUser() {
	cmd := exec.Command("git", "config", "--global", "user.name", "luisinhodamassa")

	out, _ := cmd.Output()

	fmt.Println(string(out))

}

func SetGitUserCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "git [context]",
		Short: "Sets the git user and authentication",
		Long:  "Sets the git user for the provided context according to the given scope (currently only personal and work are supported)",
		Run: func(cmd *cobra.Command, args []string) {
			setGitUser()
		},
	}
}
