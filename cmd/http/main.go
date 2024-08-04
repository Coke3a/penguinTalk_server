package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/Coke3a/TalkPenguin/internal/adapter/config"
	"github.com/Coke3a/TalkPenguin/internal/adapter/handler/http"
	"github.com/Coke3a/TalkPenguin/internal/adapter/storage/postgres"
	"github.com/Coke3a/TalkPenguin/internal/adapter/storage/postgres/repository"
	"github.com/Coke3a/TalkPenguin/internal/core/service"
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
	messageHandler := http.NewMessageHandler(messageService)

	conversationService := service.NewConversationService(conversationRepository, messageService, promptService)
	conversationHandler := http.NewConversationHandler(conversationService)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := http.NewUserHandler(userService)

	// Init router
	router, err := http.NewRouter(
		config.HTTP,
		*messageHandler,
		*conversationHandler,
		*userHandler,
	)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	// Start server
	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.URL, config.HTTP.Port)
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	err = router.Serve(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}

}
