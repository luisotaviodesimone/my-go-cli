package cmds

import (
	"encoding/base64"
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

func decodeBase64(code string) string {

	rawDecodedText, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		panic(err)
	}

	clipboard.WriteAll(string(rawDecodedText))

	fmt.Printf("Decoded text: %s\n", rawDecodedText)
	return ""
}

func Decode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decode [code]",
		Short: "retorna o código decodificado",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			decodeBase64(args[0])

		},
	}

	return cmd
}
