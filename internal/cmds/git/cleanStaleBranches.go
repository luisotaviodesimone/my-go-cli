package lodsgit

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/luisotaviodesimone/my-go-cli/internal/constants"

	"github.com/spf13/cobra"
)

func checkIfError(err error) {

	if err == nil {
		return
	}

	fmt.Printf("error: %s", err)
	os.Exit(1)
}

func logInfo(messages ...string) {
	message := ""
	for _, m := range messages {
		message += m
	}
	fmt.Printf("%s%s%s\n", constants.Blue, message, constants.Reset)
}
func logFail(messages ...string) {
	message := ""
	for _, m := range messages {
		message += m
	}
	fmt.Printf("%s%s%s\n", constants.Red, message, constants.Reset)
}

func CleanStaleBranches() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "branch-clean",
		Short: "limpar branches locais que não estão remotas",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {

			path := "./"
			repo, err := git.PlainOpen(path)
			checkIfError(err)

			ref, err := repo.Branches()
			checkIfError(err)

			staleBranches := make(map[plumbing.ReferenceName]plumbing.Hash)
			logInfo("Stale branches: ")
			err = ref.ForEach(func(r *plumbing.Reference) error {
				fmt.Printf("local: %v\n", r.Name().Short())

				config, err := repo.Config()
				if err != nil {
					return err
				}

				branchConfig, ok := config.Branches[r.Name().Short()]
				if !ok || branchConfig.Remote == "" || branchConfig.Merge == "" {
					logInfo("  No upstream configured")
					return nil
				}

				upstreamRefName := plumbing.NewRemoteReferenceName(branchConfig.Remote, branchConfig.Merge.Short())
				_, err = repo.Reference(upstreamRefName, true)
				if err != nil {
					staleBranches[r.Name()] = r.Hash()
					logFail("  Failed to find upstream reference: ", err.Error())
					return nil
				}

				for name, hash := range staleBranches {
					message := fmt.Sprintf("  %s: %s", name, hash)
					logFail(message)
				}

				return nil
			})

			checkIfError(err)
		},
	}

	return cmd
}
