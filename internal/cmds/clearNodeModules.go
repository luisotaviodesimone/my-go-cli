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

func clearNodeModules(baseDir string) {

	const (
		folderToClear = "node_modules"
	)

	fileInfo, _ := os.Stat(baseDir)

	if !fileInfo.IsDir() {
		return
	}

	curDirFiles, _ := os.ReadDir(baseDir)

	for _, element := range curDirFiles {
		elementFullPath := fmt.Sprintf("%s\\%s", baseDir, element.Name())

		if element.Name() == folderToClear && element.IsDir() {
			removeNodeModules(baseDir, elementFullPath)
		} else if element.IsDir() {
			clearNodeModules(elementFullPath)
		}
	}

}

func ClearNodeModulesCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "clear",
		Short: "Clears `node_modules` folder recursively from the current directory",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			baseDir, _ := os.Getwd()
			clearNodeModules(baseDir)
		},
	}
}
