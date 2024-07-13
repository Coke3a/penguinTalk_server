package http

import 	(
	"github.com/Coke3a/TalkPenguin/internal/core/port"
	"github.com/Coke3a/TalkPenguin/internal/core/domain"
	"github.com/gin-gonic/gin"
)

type ConversationHandler struct {
	svc port.ConversationService
}

func NewConversationHandler(svc port.ConversationService) *ConversationHandler {
	return &ConversationHandler{
		svc,
	}
}

type createConversationRequest struct {
	UserId     uint64 `json:"user_id" binding:"required" example:"1"`
	PromptId    uint64 `json:"prompt_id" binding:"required,email" example:"1"`
}

// Register godoc
//
//	@Summary		Register a new user
//	@Description	create a new user account with default role "cashier"
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			registerRequest	body		registerRequest	true	"Register request"
//	@Success		200				{object}	userResponse	"User created"
//	@Failure		400				{object}	errorResponse	"Validation error"
//	@Failure		401				{object}	errorResponse	"Unauthorized error"
//	@Failure		404				{object}	errorResponse	"Data not found error"
//	@Failure		409				{object}	errorResponse	"Data conflict error"
//	@Failure		500				{object}	errorResponse	"Internal server error"
//	@Router			/users [post]
func (ch *ConversationHandler) CreateConversation(ctx *gin.Context) {
	var req createConversationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	conversation := domain.Conversation{
		UserId:     req.UserId,
		PromptId:    req.PromptId,
	}

	_, err := ch.svc.CreateConversation(ctx, &conversation)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newCreateConversationResponse(&conversation)

	handleSuccess(ctx, rsp)
}


func (ch *ConversationHandler) CreateConversationWithMessage(ctx *gin.Context) {
	var req createConversationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	conversation := domain.Conversation{
		UserId:     req.UserId,
		PromptId:    req.PromptId,
	}

	_, message, err := ch.svc.CreateConversationWithMessage(ctx, &conversation)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newCreateConversationWithMessageResponse(&conversation, message)

	handleSuccess(ctx, rsp)
}
