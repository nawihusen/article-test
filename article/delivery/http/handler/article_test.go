package handler

import (
	"alpha-test/domain"
	"alpha-test/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*
untuk unit test pada handler tidak bisa langsung dijalankan semua menggunakan go test ./.. dan hanya bisa di jalankan
satu per satu dengan mengklik run test yang ada pada bagian atas t.run
*/

func TestPostArticle(t *testing.T) {
	UsecaseMock := new(mocks.ArticleUsecase)

	reqBody := domain.Article{
		Author: "Nagatsuki",
		Title:  "Re-Zero",
		Body:   "An Adventure",
	}

	reqBodyNotComplate := domain.Article{
		Author: "Nagatsuki",
		Body:   "An Adventure",
	}

	jBody, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal("ERROR")
	}

	NCBody, err := json.Marshal(reqBodyNotComplate)
	if err != nil {
		t.Fatal("ERROR")
	}

	app := fiber.New()
	handler := ArticleHandler{ArticleUsecase: UsecaseMock}
	basePath := viper.GetString("server.base_path")
	path := app.Group(basePath)

	path.Post("/article", handler.PostArticleHandler)

	t.Run("Success Post Article", func(t *testing.T) {
		UsecaseMock.On("PostArticle", mock.Anything, mock.Anything).Return(nil).Once()

		req := httptest.NewRequest("POST", "/article", bytes.NewBuffer(jBody))
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

		resp, err := app.Test(req)
		if err != nil {
			t.Fatal("ERROR")
		}

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NoError(t, err)
	})

	t.Run("Fail - Some Error Occured on Post Article", func(t *testing.T) {
		UsecaseMock.On("PostArticle", mock.Anything, mock.Anything).Return(errors.New("database error")).Once()

		req := httptest.NewRequest("POST", "/article", bytes.NewBuffer(jBody))
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatal("ERROR")
		}

		assert.Equal(t, 500, resp.StatusCode)
		assert.NoError(t, err)
	})

	t.Run("Fail - Client do not give all required data", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/article", bytes.NewBuffer(NCBody))
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

		resp, err := app.Test(req)
		if err != nil {
			t.Fatal("ERROR")
		}

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.NoError(t, err)
	})

	t.Run("Fail - Body parser error", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/article", nil)
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

		resp, err := app.Test(req)
		if err != nil {
			t.Fatal("ERROR")
		}

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.NoError(t, err)
	})
}

func TestGetArticle(t *testing.T) {
	novel := []domain.Article{
		{Author: "Nagatsuki", Title: "Re-Zero", Body: "An Adventure"},
		{Author: "Kinugasa", Title: "Youjitsu", Body: "A Romance"},
		{Author: "Izushiro", Title: "Saikyou Mahoushi", Body: "A Magic Adventure"},
	}

	UsecaseMock := new(mocks.ArticleUsecase)

	app := fiber.New()
	handler := ArticleHandler{ArticleUsecase: UsecaseMock}
	basePath := viper.GetString("server.base_path")
	path := app.Group(basePath)

	path.Get("/article", handler.GetArticlesHandler)

	t.Run("Success Get Article", func(t *testing.T) {
		UsecaseMock.On("GetArticles", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(novel, nil).Once()

		req := httptest.NewRequest("GET", "/article", nil)

		resp, err := app.Test(req)
		if err != nil {
			t.Fatal("ERROR")
		}

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NoError(t, err)
	})

	t.Run("Fail - an Error Occured", func(t *testing.T) {
		UsecaseMock.On("GetArticles", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("logic error")).Once()

		req := httptest.NewRequest("GET", "/article", nil)

		resp, err := app.Test(req)
		if err != nil {
			t.Fatal("ERROR")
		}

		assert.Equal(t, 500, resp.StatusCode)
		assert.NoError(t, err)
	})

	t.Run("Success but no data found", func(t *testing.T) {
		UsecaseMock.On("GetArticles", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]domain.Article{}, nil).Once()

		req := httptest.NewRequest("GET", "/article", nil)

		resp, err := app.Test(req)
		if err != nil {
			t.Fatal("ERROR")
		}

		assert.Equal(t, 404, resp.StatusCode)
		assert.NoError(t, err)
	})

}
