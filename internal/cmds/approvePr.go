package cmds

import (
	"context"

	"github.com/cli/go-gh/v2"
	"github.com/spf13/cobra"
)

func approvePr(addGif bool) {
	body := ""

	if addGif {
		body = "![gif](https://i.shipit.today/)"
	}

	ctx := context.Background()

	gh.ExecInteractive(ctx, "pr", "review", "--approve", "-b", body)

}

func ApprovePrCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve",
		Short: "Approves current branch's PR",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			shouldUseGif, err := cmd.Flags().GetBool("add-gif")

			if err != nil {
				panic(err)
			}

			approvePr(shouldUseGif)
		},
	}

	cmd.PersistentFlags().BoolP("add-gif", "g", false, "Flag to indicate if the gif should be added to the PR body")

	return cmd
}
