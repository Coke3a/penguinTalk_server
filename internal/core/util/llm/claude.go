package llm

import (
	// "bytes"
	// "encoding/json"
	// "io/ioutil"
	// "net/http"
	// "os"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	enum "github.com/Coke3a/TalkPenguin/internal/adapter/enum"
	"github.com/Coke3a/TalkPenguin/internal/core/domain"
)

// Enum for MessageType
type MessageType int

const (
	User MessageType = iota
	AI
)

// Structs for request payload
type Content struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

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

// Role function to get the role based on MessageType
func Role(msgType enum.MessageType) string {
	if msgType == enum.MessageTypes.User {
		return "user"
	}
	return "assistant"
}

func RequestGenerateMessage(prompt string, prompt2 string, messages []domain.Message) (string, error) {
    // Construct messages for the payload
    var messagePayloads []MessagePayload

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
    } else {
        // Add default "empty" message if messages is nil or empty
        messagePayloads = append(messagePayloads, MessagePayload{
            Role: "user",
            Content: []Content{
                {
                    Type: "text",
                    Text: "‘empty’",
                },
            },
        })
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
        return "", err
    }
    body := bytes.NewReader(payloadBytes)

    // Create the HTTP request
    req, err := http.NewRequest("POST", "https://api.anthropic.com/v1/messages", body)
    if err != nil {
        return "", err
    }
    // req.Header.Set("X-Api-Key", os.ExpandEnv("$ANTHROPIC_API_KEY"))
	req.Header.Set("X-Api-Key", "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")
    req.Header.Set("Anthropic-Version", "2023-06-01")
    req.Header.Set("Content-Type", "application/json")

    // Send the request
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    // Read the response
    responseBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return string(responseBytes), nil
}