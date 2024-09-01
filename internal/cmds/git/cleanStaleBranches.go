package lodsgit

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/plumbing/cache"
	"github.com/go-git/go-git/storage/filesystem"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"

	. "github.com/go-git/go-git/v5/_examples"
	// colors "github.com/luisotaviodesimone/my-go-cli/internal/constants"
	// "github.com/luisotaviodesimone/my-go-cli/internal/utils"
	"github.com/spf13/cobra"
)

func CleanStaleBranches() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "branch-clean",
		Short: "limpar branches locais que não estão remotas",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			Info("git log --oneline -1")

			path := "./"
			fs := osfs.New(path)

			s := filesystem.NewStorageWithOptions(fs, cache.NewObjectLRUDefault(), filesystem.Options{KeepDescriptors: true})
			r, err := git.Open(s, fs)
			CheckIfError(err)
			defer s.Close()

			revision := "--is-inside-work-tree"
			h, err := r.ResolveRevision(plumbing.Revision(revision))
			CheckIfError(err)

			commit, err := r.CommitObject(*h)

			commitIter, err := r.Log(&git.LogOptions{From: commit.Hash})
			CheckIfError(err)

			err = commitIter.ForEach(func(c *object.Commit) error {
				hash := c.Hash.String()
				line := strings.Split(c.Message, "\n")
				fmt.Println(hash[:7], line[0])

				return nil
			})
			CheckIfError(err)

			log.Println("git branch-clean called")
		},
	}

	return cmd
}
