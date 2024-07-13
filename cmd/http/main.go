package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/Coke3a/TalkPenguin/internal/adapter/config"
	"github.com/Coke3a/TalkPenguin/internal/adapter/storage/postgres"
	"github.com/Coke3a/TalkPenguin/internal/adapter/storage/postgres/repository"
	"github.com/Coke3a/TalkPenguin/internal/core/service"

	"github.com/Coke3a/TalkPenguin/internal/core/domain"
)

func main() {
	// Load environment variables
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	slog.Info("Starting the application", "app", config.App.Name, "env", config.App.Env)

	// Init database
	ctx := context.Background()
	db, err := postgres.Connect(ctx, config.DB)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	slog.Info("Successfully connected to the database", "db", config.DB.Connection)

	// Migrate database
	err = db.Migrate()
	if err != nil {
		slog.Error("Error migrating database", "error", err)
		os.Exit(1)
	}

	slog.Info("Successfully migrated the database")

	modelTransactionRepository := repository.NewModelTransactionRepository(db)
	modelTransactionService := service.NewModelTransactionService(modelTransactionRepository)

	conversationRepository := repository.NewConversationRepository(db)
	
	promptRepository := repository.NewPromptRepository(db)
	promptService := service.NewPromptService(promptRepository)

	messageRepository := repository.NewMessageRepository(db)
	messageService := service.NewMessageService(messageRepository, conversationRepository, promptService, modelTransactionService)

	// conversationService := service.NewConversationService(conversationRepository, messageService, promptService)
	
	// conversation := domain.Conversation{
	// 	UserId: 1,
	// 	PromptId: 1,
	// }

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	
	// conversationService.CreateConversation(ctx, &conversation)
	// _, message, err := conversationService.CreateConversationWithMessage(ctx, &conversation)
	// if (err != nil ) {
	// 	slog.Error("Error creating conversation", "error", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(message.MessageText)
	// fmt.Println("Hello, World!")



	
	message := &domain.Message{
		ConversationId: 1,
		UserId: 1,
		MessageText: "I think I'll go to write the daily memo.",
		MessageAudio: "123456",
	}

	message, err = messageService.ExchangingMessage(ctx, message)
	if (err != nil ) {
		slog.Error("Error creating conversation", "error", err)
		os.Exit(1)
	}
	fmt.Println(message.MessageText)
	fmt.Println("Hello, World!")
	// user_id, conversation_id message_text message_audio
}
