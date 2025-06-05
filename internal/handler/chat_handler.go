package handler

import (
	"github.com/crazydw4rf/porto-chatbot/internal/instruction"
	"github.com/crazydw4rf/porto-chatbot/internal/services"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/genai"
)

type ChatHandler struct {
	ChatServices services.ChatServices
}

func NewChatHandler(chatServices services.ChatServices) *ChatHandler {
	return &ChatHandler{chatServices}
}

func (h *ChatHandler) AskPorto(ctx *fiber.Ctx) error {
	q := new(struct{ Prompt string })

	err := ctx.BodyParser(q)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	h.ChatServices.SetSystemInstructionText(instruction.UcupPortfolio, genai.RoleModel)

	response, err := h.ChatServices.GenerateContent(ctx.Context(), q.Prompt)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate content",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"response": response})
}
