package cmds

import (
	"fmt"
	"os"
	"strings"

	colors "github.com/luisotaviodesimone/my-go-cli/internal/constants"
	"github.com/spf13/cobra"
)

func removeNodeModules(parentFolder string, nodeModulesFullPath string) {
	parentFolderArray := strings.Split(parentFolder, "\\")
	parentFolderName := parentFolderArray[len(parentFolderArray)-1]

	fmt.Printf("Removing %snode_modules%s folder from %s%s%s\n", colors.Red, colors.Reset, colors.Green, parentFolderName, colors.Reset)

	os.RemoveAll(nodeModulesFullPath)

	fmt.Printf("\n%s%s%s\n\n", colors.Green, "Done!", colors.Reset)
}

func clearNodeModules(baseFolder string, shouldNotRecurse bool) {

	const (
		folderToClear = "node_modules"
	)

	fileInfo, _ := os.Stat(baseFolder)

	if !fileInfo.IsDir() {
		return
	}

	curDirFiles, _ := os.ReadDir(baseFolder)

	for _, element := range curDirFiles {
		elementFullPath := fmt.Sprintf("%s/%s", baseFolder, element.Name())

		if element.Name() == folderToClear && element.IsDir() {
			removeNodeModules(baseFolder, elementFullPath)
		} else if element.IsDir() && !shouldNotRecurse {
			clearNodeModules(elementFullPath, shouldNotRecurse)
		}
	}

}

func ClearNodeModulesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear",
		Short: "Clears `node_modules` folder from the current directory and its subdirectories",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {

			var baseFolder string
			var err error

			baseFolder = cmd.Flag("base-folder").Value.String()
			shouldNotRecurse, _ := cmd.Flags().GetBool("single-execution")

			if baseFolder == "" {
				baseFolder, err = os.Getwd()
			}

			if err != nil {
				fmt.Println("Error getting current directory")
				return
			}

			clearNodeModules(baseFolder, shouldNotRecurse)
		},
	}

	cmd.PersistentFlags().StringP("base-folder", "f", "", "Folder to start the cleaning from")
	cmd.PersistentFlags().BoolP("single-execution", "s", false, "If the cleaning should only be done in the current directory")

	return cmd
}
