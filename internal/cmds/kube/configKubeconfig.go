package kube

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/luisotaviodesimone/my-go-cli/internal/colors"
	"github.com/luisotaviodesimone/my-go-cli/internal/utils"
	"github.com/spf13/cobra"
)

func ensureFolder(path string) []os.DirEntry {
	folder, err := os.ReadDir(path)

	if err != nil {

		fmt.Printf("%sProblema ao abrir o diretório %s\n%s\n%s", colors.Red, path, err, colors.Reset)
		fmt.Printf("\nCriando o diretório %s%s%s\n", colors.Blue, path, colors.Reset)

		os.MkdirAll(path, 0777)

	}

	return folder
}

func replaceAndWriteConfigFile(
	folderPath string,
	newFileAndContextName string,
	newIpAddress string,
	oldFileAndContextName string,
	oldIpAddress string,
) {

	filePath := filepath.Join(folderPath, newFileAndContextName)

	file, _ := os.Open(filePath)

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		newLine := strings.ReplaceAll(
			strings.ReplaceAll(scanner.Text(), oldFileAndContextName, newFileAndContextName),
			oldIpAddress,
			newIpAddress,
		)

		lines = append(lines, newLine)
	}

	error := os.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0777)

	if error != nil {
		log.Fatalf("Erro manipulando o arquivo\n%s%s%s\n", colors.Red, error, colors.Reset)
	}
}

func setKubeconfigFile(
	newFileAndContextName string,
	newIpAddress string,
	oldFileAndContextName string,
	oldIpAddress string,
) {

	kubeconfigsFolderPath := filepath.Join(os.Getenv("HOME"), ".kube", "configs")

	folder := ensureFolder(kubeconfigsFolderPath)

	for _, e := range folder {

		if e.Name() == newFileAndContextName {

			replaceAndWriteConfigFile(
				kubeconfigsFolderPath,
				newFileAndContextName,
				newIpAddress,
				oldFileAndContextName,
				oldIpAddress,
			)

			return
		}
	}

	log.Fatalf("%sErro na configurando o arquivo!%s", colors.Red, colors.Reset)

}

func ConfigKubeConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "manipulação da configurção do `kubectl`",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			var newContext string
			var newIp string

			newContext = cmd.Flag("context").Value.String()
			newIp = cmd.Flag("ip").Value.String()

			if newContext == "" {
				wordPromptContent := utils.PromptContent{
					ErrorMsg: "Forneça o nome/contexto do arquivo de `kubeconfig`.",
					Label:    "Nome/Contexto do `kubeconfig`: ",
				}

				newContext = utils.PromptGetInput(wordPromptContent)
			}

			if newIp == "" {
				wordPromptContent := utils.PromptContent{
					ErrorMsg: "Forneça o ip para o arquivo de `kubeconfig`.",
					Label:    "Ip para o `kubeconfig`?",
				}

				newIp = utils.PromptGetInput(wordPromptContent)
			}

			setKubeconfigFile(newContext, newIp, "default", "127.0.0.1")
		},
	}

	cmd.PersistentFlags().StringP("context", "c", "", "Context name to be substituted for")
	cmd.PersistentFlags().StringP("ip", "i", "", "Ip address to be substituted for")

	return cmd
}
