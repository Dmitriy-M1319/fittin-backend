package llm

import (
	"fmt"

	"github.com/Dmitriy-M1319/fittin-backend/internal/mmpi/config"
	"github.com/Dmitriy-M1319/fittin-backend/internal/mmpi/models"
	"github.com/gage-technologies/mistral-go"
)

type LLMClient struct {
	client *mistral.MistralClient
}

func NewLLMClient(cfg config.Config) *LLMClient {
	return &LLMClient{client: mistral.NewMistralClientDefault(cfg.Llm.ApiKey)}
}

func (cl *LLMClient) PrepareTestResult(result models.TestResult) (string, error) {
	scaleValues := make(map[string]int)
	for _, scale := range result.Scales {
		scaleValues[scale.Scale] = scale.Value
	}
	request := fmt.Sprintf(requestFormatString, scaleValues["(Hs)"],
		scaleValues["(D)"],
		scaleValues["(Hy)"],
		scaleValues["(Pd)"],
		scaleValues["(Mf)"],
		scaleValues["(Pa)"],
		scaleValues["(Pt)"],
		scaleValues["(Sc)"],
		scaleValues["(Ma)"],
		scaleValues["(Si)"])

	chatRes, err := cl.client.Chat("mistral-large-latest",
		[]mistral.ChatMessage{{Content: request, Role: mistral.RoleUser}}, nil)
	return chatRes.Choices[0].Message.Content, err
}
