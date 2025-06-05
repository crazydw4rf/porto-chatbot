package services

import (
	"io/fs"
	"net/http"

	"github.com/crazydw4rf/porto-chatbot/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

type FiberService struct {
	App *fiber.App
}

func NewFiberService(cfg *config.Config) FiberService {
	app := fiber.New()

	app.Use(cors.New(cors.Config{AllowOrigins: cfg.CORS_ORIGINS}))

	return FiberService{app}
}

func (f FiberService) SetStaticEmbeddedFiles(path string, prefix string, fs fs.FS) {
	f.App.Use("/", filesystem.New(filesystem.Config{
		Browse:     false,
		Root:       http.FS(fs),
		PathPrefix: prefix,
	}))
}
