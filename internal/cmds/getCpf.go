package cmds

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

func buildRequest(state string) *http.Request {
	formValuesMap := map[string]string{
		"acao":       "gerar_cpf",
		"pontuacao":  "S",
		"cpf_estado": state,
	}

	var formValuesString strings.Builder

	for key, element := range formValuesMap {
		fmt.Fprintf(&formValuesString, "%s=%s&", key, element)
	}

	payload := strings.NewReader(formValuesString.String())
	url := "https://www.4devs.com.br/ferramentas_online.php"

	newRequest, err := http.NewRequest(http.MethodPost, url, payload)

	if err != nil {
		panic(err)
	}

	newRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	return newRequest
}

func makeRequest(state string) (string, error) {

	req := buildRequest(state)

	clientRes, clientErr := http.DefaultClient.Do(req)

	if clientErr != nil {
		return "", clientErr
	}

	defer clientRes.Body.Close()

	clientResBody, readErr := io.ReadAll(clientRes.Body)

	if readErr != nil {
		return "", readErr
	}

	return string(clientResBody), nil
}

func GetCpfCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cpf",
		Short: "retorna um cpf válido",
		Run: func(cmd *cobra.Command, args []string) {
			stateFlag := cmd.Flag("state").Value.String()

			cpf, err := makeRequest(stateFlag)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			clipboard.WriteAll(cpf)
			fmt.Print(cpf)
		},
	}

	cmd.PersistentFlags().StringP("state", "s", "AC", "The state for the CPF to be generated")

	return cmd
}
