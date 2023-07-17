package cmds

import (
	"fmt"
	"os"

	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/handlers"
	"github.com/luisotaviodesimone/my-go-cli/internal/utils"
	"github.com/spf13/cobra"
)

func speak(text string, voice string) {
	speech := htgotts.Speech{Folder: "audio", Language: voice, Handler: &handlers.Native{}}
	speech.Speak(text)

	currentDir := utils.GetCurrentExecutableDirPath()

	dirPath := fmt.Sprintf("%s/%s", currentDir, "audio")
	audioDir, _ := os.ReadDir(dirPath)

	for _, element := range audioDir {
		os.Remove(fmt.Sprintf("%s/%s", dirPath, element.Name()))
	}

}

func Speak() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "speak [text]",
		Short: "Speaks the given text",
		Long:  "",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			voiceFlag := cmd.Flag("voice").Value.String()

			speak(args[0], voiceFlag)
		},
	}

	cmd.PersistentFlags().StringP("voice", "v", "pt-br", "The voice to be used")

	return cmd

}
