package domain

import (
	"time"

	enum "github.com/Coke3a/TalkPenguin/internal/adapter/enum"
)

type Message struct {
	MessageId      uint64
	ConversationId uint64
	UserId         uint64
	MtId           uint64
	MessageText    string
	MessageAudio   string
	MessageType    enum.MessageType
	MessageDate    *time.Time
}
