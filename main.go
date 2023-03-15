package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func GetAnswer(message string) (string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{Role: openai.ChatMessageRoleSystem, Content: "You are a succinct and specific mentor"},
				{Role: openai.ChatMessageRoleUser, Content: message},
			},
			Temperature: 0.5,
			N:           1,
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, err
}

func main() {
	exePath, _ := os.Executable()

	// Carrega o .env do diretorio do executavel
	godotenv.Load(filepath.Dir(exePath) + "\\.env")
	// Carrega o .env do diretorio corrente
	godotenv.Load()

	key := os.Getenv("OPENAI_API_KEY")

	if key == "" {
		fmt.Println("A variável de ambiente OPENAI_API_KEY precisa conter um token válido da OpenAI")
		os.Exit(2)
	}

	if len(os.Args) < 2 {
		fmt.Println("Uso: ask <texto a ser processado pelo gpt-3.5>")
		os.Exit(1)
	}

	message := strings.Join(os.Args[1:], " ")

	answer, err := GetAnswer(message)

	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		os.Exit(3)
	}

	fmt.Println(answer)
}
