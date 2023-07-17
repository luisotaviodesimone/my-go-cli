package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func GetFolderSize(folderPath string) (int64, error) {
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
	return size, nil
}

func main() {

	out, err := exec.Command("scoop", "list").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", string(out))

}
