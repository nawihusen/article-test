package http

import (
	handler "alpha-test/article/delivery/http/handler"
	"alpha-test/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// RouterAPI is main router for this Service Saksi REST API
func RouterAPI(app *fiber.App, article domain.ArticleUsecase) {
	articleHandler := &handler.ArticleHandler{ArticleUsecase: article}

	basePath := viper.GetString("server.base_path")
	path := app.Group(basePath)

	// Article Management
	path.Post("/article", articleHandler.PostArticle)
	path.Get("/article", articleHandler.GetArticles)
}
