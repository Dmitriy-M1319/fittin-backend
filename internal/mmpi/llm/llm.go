package llm

import (
	"github.com/Dmitriy-M1319/fittin-backend/internal/mmpi/models"
	"github.com/gage-technologies/mistral-go"
)

type LLMClient struct {
	client *mistral.MistralClient
}

func NewLLMClient() *LLMClient {
	return &LLMClient{client: mistral.NewMistralClientDefault("api-key")}
}

func (cl *LLMClient) PrepareTestResult(result models.TestResult) (string, error) {

	chatRes, err := cl.client.Chat("mistral-large-latest",
		[]mistral.ChatMessage{{Content: "Hello, world!", Role: mistral.RoleUser}}, nil)
	return chatRes.Object, err
}
