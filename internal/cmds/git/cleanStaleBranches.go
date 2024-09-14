package lodsgit

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"

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

			sshKeyFilePath := fmt.Sprintf("%s/.ssh/%s", os.Getenv("HOME"), "id_rsa")
			keyFileBytes, err := os.ReadFile(sshKeyFilePath)

			if err != nil {
				log.Fatal("Unable to open ssh key file: ", err)
			}

			publicKeys, err := ssh.NewPublicKeys("git", keyFileBytes, "")
			if err != nil {
				log.Fatalf("Failed to create public keys: %v", err)
			}

			err = repo.Fetch(&git.FetchOptions{
				RemoteName: "origin",
				Auth:       publicKeys,
				Progress:   os.Stdout,
				Prune:      true,
			})
			if err != nil && err != git.NoErrAlreadyUpToDate {
				fmt.Println("Fetch failed: ", err)
			}

			ref, err := repo.Branches()

			staleBranches := make(map[plumbing.ReferenceName]plumbing.Hash)
			logInfo("Stale branches: ")
			err = ref.ForEach(func(r *plumbing.Reference) error {
				config, err := repo.Config()
				if err != nil {
					return err
				}

				branchConfig, _ := config.Branches[r.Name().Short()]

				upstreamRefName := plumbing.NewRemoteReferenceName(branchConfig.Remote, branchConfig.Merge.Short())
				_, err = repo.Reference(upstreamRefName, true)
				if err != nil {
					staleBranches[r.Name()] = r.Hash()
					return nil
				}

				return nil
			})

			for name := range staleBranches {
				message := fmt.Sprintf("  %s", name)
				logInfo(message)
			}

			checkIfError(err)
		},
	}

	return cmd
}
