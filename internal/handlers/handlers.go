package handler

import (
	"context"
	"mime/multipart"

	"project_sem/internal/model"

	"github.com/gofiber/fiber/v2"
)

type MarketingService interface {
	SaveProducts(ctx context.Context, file *multipart.FileHeader) (*model.LoadResult, error)
}

type MarketingHandler struct {
	app     *fiber.App
	service MarketingService
}

func RegisterRoutes(
	app *fiber.App,
	service MarketingService,
) *MarketingHandler {
	handler := &MarketingHandler{
		app:     app,
		service: service,
	}

	handler.app.Post("/api/v0/prices", handler.UploadProducts)

	return handler
}