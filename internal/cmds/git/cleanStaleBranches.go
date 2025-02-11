package lodsgit

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"

	"github.com/luisotaviodesimone/my-go-cli/internal/colors"

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
	fmt.Printf("%s%s%s\n", colors.Blue, message, colors.Reset)
}
func logFail(messages ...string) {
	message := ""
	for _, m := range messages {
		message += m
	}
	fmt.Printf("%s%s%s\n", colors.Red, message, colors.Reset)
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

			checkIfError(err)

			for name := range staleBranches {
				message := fmt.Sprintf("  %s", name)
				logInfo(message)
			}

			reader := bufio.NewReader(os.Stdin)
			for {
				fmt.Print("Remove references (Y/n): ")
				text, error := reader.ReadString('\n')

				if error != nil {
					log.Fatalf("error reading from given string: %s", error)
				}
				lowerText := strings.ToLower(text)

				if lowerText == "y\n" || lowerText == "yes\n" || lowerText == "\n" {

					for name := range staleBranches {
						repo.Storer.RemoveReference(name)
						fmt.Println("Removed", colors.Cyan, name, colors.Reset)
					}

					break
				} else if lowerText == "n\n" || lowerText == "no\n" {
					fmt.Println("exiting...")
					break
				}
			}
		},
	}

	return cmd
}
