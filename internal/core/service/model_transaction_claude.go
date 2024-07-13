package service

import (
	"context"
	"errors"

	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Coke3a/TalkPenguin/internal/core/domain"
	"github.com/Coke3a/TalkPenguin/internal/core/port"

	enum "github.com/Coke3a/TalkPenguin/internal/adapter/enum"
)

type ModelTransactionService struct {
	repo  port.ModelTransactionRepository
	// cache port.CacheRepository
}

func NewModelTransactionService(repo port.ModelTransactionRepository) *ModelTransactionService {
	return &ModelTransactionService{
		repo,
		// cache,
	}
}


// Enum for MessageType
type MessageType int

const (
	User MessageType = iota
	AI
)

type MessagePayload struct {
	Role    string    `json:"role"`
	Content []Content `json:"content"`
}

type Payload struct {
	Model     string           `json:"model"`
	MaxTokens int              `json:"max_tokens"`
	System    string           `json:"system"`
	Messages  []MessagePayload `json:"messages"`
}

type Content struct {
    Type string `json:"type"`
    Text string `json:"text"`
}

type Usage struct {
    InputTokens  int `json:"input_tokens"`
    OutputTokens int `json:"output_tokens"`
}

// Define the response structure
type Response struct {
    ID           string    `json:"id"`
    Type         string    `json:"type"`
    Role         string    `json:"role"`
    Model        string    `json:"model"`
    Content      []Content `json:"content"`
    StopReason   string    `json:"stop_reason"`
    StopSequence *string   `json:"stop_sequence"`
    Usage        Usage     `json:"usage"`
}



// Role function to get the role based on MessageType
func Role(msgType enum.MessageType) string {
	if msgType == enum.MessageTypes.User {
		return "user"
	}
	return "assistant"
}

func (mts *ModelTransactionService) RequestToModel(ctx context.Context, prompt string, prompt2 string, messages []domain.Message) (*domain.ModelTransaction, error) {

    // Construct messages for the payload
    var messagePayloads []MessagePayload
    
    messagePayloads = append(messagePayloads, MessagePayload{
        Role: "user",
        Content: []Content{
            {
                Type: "text",
                Text: "‘empty’",
            },
        },
    })

    if len(messages) > 0 {
        for _, msg := range messages {
            messagePayloads = append(messagePayloads, MessagePayload{
                Role: Role(msg.MessageType),
                Content: []Content{
                    {
                        Type: "text",
                        Text: msg.MessageText,
                    },
                },
            })
        }
    }

    // Prepare the payload
    data := Payload{
        Model:     "claude-3-haiku-20240307",
        MaxTokens: 1024,
        System:    prompt + " " + prompt2,
        Messages:  messagePayloads,
    }

    // Marshal the payload to JSON
    payloadBytes, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }
    body := bytes.NewReader(payloadBytes)

    // Create the HTTP request
    req, err := http.NewRequest("POST", "https://api.anthropic.com/v1/messages", body)
    if err != nil {
        return nil, err
    }
    // req.Header.Set("X-Api-Key", os.ExpandEnv("$ANTHROPIC_API_KEY"))
	req.Header.Set("X-Api-Key", "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")
    req.Header.Set("Anthropic-Version", "2023-06-01")
    req.Header.Set("Content-Type", "application/json")

    // Send the request
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // Read the response
    responseBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

	modelTransaction, err := mts.repo.CreateModelTransaction(ctx, payloadBytes, responseBytes)
	if err != nil {
        return nil, err
    }
	
    return modelTransaction, nil
}

func (mts *ModelTransactionService) GetResponseMessage(ctx context.Context, modelTransaction *domain.ModelTransaction) (string, error) {

	// Unmarshal JSON data into Response struct
	var response Response
	err := json.Unmarshal([]byte(modelTransaction.ResponseData), &response)
	if err != nil {
        return "", err
    }

	// Extract the text content
	for _, content := range response.Content {
		if content.Type == "text" {
			return content.Text, nil
		}
	}
	
	return "", 	errors.New(`Response text message not found`)
}	