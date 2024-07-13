package domain

type Prompt struct {
	PromptId            uint64
	ConversationTopicId uint64
	PromptLangId        uint64
	Prompt              string
	Prompt2             string
	AiRole              uint64
	UserRole            uint64
}
