package domain

import (
	"time"
)

type Conversation struct {
	ConversationId    uint64
	UserId            uint64
	PromptId          uint64
	ConversationStart *time.Time
	ConversationEnd   *time.Time
}
