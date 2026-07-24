package internal

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
	"google.golang.org/genai"
)

func getAPIKey() string {
	data, err := os.ReadFile(".config")
	if err == nil {
		key := strings.TrimSpace(string(data))
		if key != "" {
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
		APIKey: "<paste_apikey_here>",
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
