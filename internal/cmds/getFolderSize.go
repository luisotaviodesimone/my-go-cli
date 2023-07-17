package cmds

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func getFolderSize(folderPath string) (float64, error) {
	var size int64
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return float64(size) / float64(MB), nil
}

func FolderSizeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "size [folder]",
		Short: "retorna o tamanho do diretório/arquivo",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			size, err := getFolderSize(args[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("O tamanho do diretório `%s` é %.4f megabytes\n", args[0], size)
		},
	}
}
