package services

import (
	"context"

	"github.com/crazydw4rf/porto-chatbot/internal/config"
	"google.golang.org/genai"
)

type ChatServices struct {
	Client *genai.Client
	Config *genai.GenerateContentConfig
}

func NewChatServices(cfg *config.Config) (ChatServices, error) {
	client, err := genai.NewClient(context.Background(), &genai.ClientConfig{
		APIKey:  cfg.GEMINI_API_KEY,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return ChatServices{}, err
	}

	return ChatServices{
		Client: client,
		Config: &genai.GenerateContentConfig{},
	}, nil
}

func (cs *ChatServices) GenerateContent(ctx context.Context, prompt string) (string, error) {
	response, err := cs.Client.Models.GenerateContent(ctx, "gemini-2.0-flash", genai.Text(prompt), cs.Config)
	if err != nil {
		return "", err
	}

	return response.Text(), nil
}

func (cs *ChatServices) SetSystemInstructionText(instruction string, role genai.Role) {
	cs.Config.SystemInstruction = genai.NewContentFromText(instruction, role)
}
