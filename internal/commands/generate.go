package commands

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	apiKey     = "96a66b32cf8e424c84632e221bbec779"
	apiVersion = "2024-02-15-preview"
	apiBase    = "https://habibiswitch.openai.azure.com"
)

func HandleGenerate(prompt string) error {
	// Commands to execute sequentially
	commands := []string{
		fmt.Sprintf("export AZURE_API_BASE=%s", apiBase),
		fmt.Sprintf("export AZURE_API_VERSION=%s", apiVersion),
		fmt.Sprintf("export AZURE_API_KEY=%s", apiKey),

		"aider --model azure/Habibi-4o",
	}

	// Execute each command sequentially
	for _, cmdStr := range commands {
		cmd := exec.Command("bash", "-c", cmdStr)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		fmt.Printf("Running command: %s\n", cmdStr)
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error running command: %s\n", err)
			return err
		}
	}

	return nil
}
