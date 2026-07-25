package internal

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/glamour"
	"google.golang.org/genai"
)

func getAPIKey() string {
	ex, err := os.Executable()
	if err != nil {
		return os.Getenv("GEMINI_API_KEY")
	}

	filePath := filepath.Join(filepath.Dir(ex), ".config")

	apiKey, err := os.ReadFile(filePath)
	if err == nil {
		key := strings.TrimSpace(string(apiKey))
		if key == "" {
			return key
		}
	}

	return os.Getenv("GEMINI_API_KEY")
}

func CallAI(inputPrompt string) {
	apiKey := getAPIKey()
	if apiKey == "" {
		log.Fatal("Error: No API Key found. run: halo set <your_api_key>")
	}

	ctx := context.Background()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	res, err := client.Models.GenerateContent(
		ctx,
		"gemini-3.5-flash-lite",
		genai.Text(inputPrompt),
		nil,
	)
	if err != nil {
		log.Fatalf("api call failed: %v", err)
	}

	out, err := glamour.Render(res.Text(), "dark")
	if err != nil {
		fmt.Println(res.Text())
		return
	}

	fmt.Print(out)
}
