package api

import (
	"embed"
	"net/http"

	"github.com/crazydw4rf/porto-chatbot/internal/config"
	"github.com/crazydw4rf/porto-chatbot/internal/handler"
	"github.com/crazydw4rf/porto-chatbot/internal/services"
	"github.com/rotisserie/eris"

	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

//go:embed web
var webFiles embed.FS

func Handler(w http.ResponseWriter, r *http.Request) {
	cfg, err := config.InitConfig()
	if err != nil {
		http.Error(w, eris.Wrap(err, "failed to initialize config").Error(), http.StatusInternalServerError)
		return
	}

	fiber_ := services.NewFiberService(cfg)
	fiber_.SetStaticEmbeddedFiles("/", "web", webFiles)

	chatServices, err := services.NewChatServices(cfg)
	if err != nil {
		http.Error(w, eris.Wrap(err, "failed to initialize chat services").Error(), http.StatusInternalServerError)
		return
	}

	chatHandler := handler.NewChatHandler(chatServices)
	fiber_.App.Post("/chat", chatHandler.AskPorto)

	r.RequestURI = r.URL.String()

	adaptor.FiberApp(fiber_.App).ServeHTTP(w, r)
}
