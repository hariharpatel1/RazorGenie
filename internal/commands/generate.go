package commands

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

var (
	apiKey     = "96a66b32cf8e424c84632e221bbec779"
	apiVersion = "2024-02-15-preview"
	apiBase    = "https://habibiswitch.openai.azure.com"
)

func HandleGenerate(dirPath string, prompt string) error {
	// Commands to execute sequentially
	commands := []string{
		fmt.Sprintf("cd %s", dirPath),
		fmt.Sprintf("export AZURE_API_BASE=%s", apiBase),
		fmt.Sprintf("export AZURE_API_VERSION=%s", apiVersion),
		fmt.Sprintf("export AZURE_API_KEY=%s", apiKey),
		fmt.Sprintf("source ~/.zshrc"),

		"aider --model azure/Habibi-4o",

		prompt,

		"/exit",
	}

	cmd := exec.Command("zsh")

	// Create a pseudo-terminal for the shell session
	ptmx, err := pty.Start(cmd)
	if err != nil {
		return fmt.Errorf("failed to start pty: %v", err)
	}
	defer func() { _ = ptmx.Close() }() // Best effort close

	// Capture the output from the pseudo-terminal
	go func() {
		_, _ = io.Copy(os.Stdout, ptmx)
	}()

	// Execute each command sequentially
	for _, cmdStr := range commands {
		_, err := ptmx.Write([]byte(cmdStr + "\n"))
		if err != nil {
			return fmt.Errorf("failed to write command: %v", err)
		}
	}

	// Wait for the shell session to complete
	if err = cmd.Wait(); err != nil {
		return fmt.Errorf("shell session exited with error: %v", err)
	}

	return nil
}
