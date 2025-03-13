package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/luisotaviodesimone/terminal-ui/spinner"
)

func main() {
	workDir, err := os.Getwd()
	s := spinner.New(spinner.Config{})

	if err != nil {
		panic(err)
	}

	cmd := exec.Command("go", "build", filepath.Join(workDir, "cmd/main.go"))

	fmt.Print("Buildando ")
	start := time.Now()
	s.Start()
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	elapsed := time.Since(start)
	s.Stop()
	fmt.Printf("\nBuild levou %s\n", elapsed.Round(time.Millisecond))

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	os.Rename(filepath.Join(workDir, "main"), filepath.Join(homeDir, ".local/bin/lods"))
}
