package cmds

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/cli/go-gh/v2"
	"github.com/luisotaviodesimone/my-go-cli/internal/colors"
	"github.com/luisotaviodesimone/my-go-cli/internal/utils"
	"github.com/spf13/cobra"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Session struct {
	User User `json:"user"`
}

func openAndParseSensibleInfoJson(sensibleInfoFilePath string) map[string]Session {
	var sessionUsers map[string]Session

	sensibleInfo, err := os.Open(sensibleInfoFilePath)

	if err != nil {
		fmt.Println("Error opening sensible-info.json file")
		return sessionUsers
	}

	defer sensibleInfo.Close()

	byteValue, _ := io.ReadAll(sensibleInfo)

	json.Unmarshal(byteValue, &sessionUsers)

	return sessionUsers
}

func setGitUser(session string, sensibleFilePath string) {
	sessionsUsers := openAndParseSensibleInfoJson(sensibleFilePath)

	ctx := context.Background()

	gh.ExecInteractive(ctx, "auth", "logout")

	loginArgs := []string{"auth", "login", "-h", "GitHub.com", "-p", "https", "-w"}

	err := gh.ExecInteractive(ctx, loginArgs...)

	if err != nil {
		fmt.Println("Login aborted")
		return
	}

	fmt.Printf(
		`
	Setting git user for %s%s%s context:
		name: %s%s%s
		email: %s%s%s

	`, colors.Purple, session, colors.Reset,
		colors.Green, sessionsUsers[session].User.Name, colors.Reset,
		colors.Green, sessionsUsers[session].User.Email, colors.Reset)

	exec.Command("git", "config", "--global", "user.name", sessionsUsers[session].User.Name).Run()
	exec.Command("git", "config", "--global", "user.email", sessionsUsers[session].User.Email).Run()

}

func SetGitUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "git [session]",
		Short: "Sets the git user and authentication",
		Long:  "Sets the git user for the provided context according to the given scope (currently only `personal` and `work` are supported)",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if args[0] != "personal" && args[0] != "work" {
				fmt.Printf("%s%s%s is not a valid session. Please use %s%s%s or %s%s%s", colors.Red, args[0], colors.Reset, colors.Green, "personal", colors.Reset, colors.Green, "work", colors.Reset)
				return
			}

			fileCredentials := cmd.Flag("file-credentials").Value.String()

			setGitUser(args[0], fileCredentials)
		},
	}

	currentDir := utils.GetCurrentExecutableDirPath()

	cmd.PersistentFlags().StringP("auth-method", "a", "https", "Authentication method to be used (ssh or https)")
	cmd.PersistentFlags().StringP("file-credentials", "f", fmt.Sprintf("%s/sensible-info.json", currentDir), "Path to the sensible info json file")

	return cmd
}
