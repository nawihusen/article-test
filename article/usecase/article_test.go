package usecase

import (
	"alpha-test/domain"
	"alpha-test/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*
gunakan go test ./article/usecase -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html untuk langsung
menjalankan semua unit test
*/

func TestPostArticle(t *testing.T) {
	mockMYSQL := new(mocks.ArticleRepository)
	mockRedis := new(mocks.ArticleRedisRepository)
	subaru := domain.Article{Author: "Nagatsuki", Title: "Re-Zero", Body: "An Adventure"}

	t.Run("Success post data to database", func(t *testing.T) {
		mockMYSQL.On("PostArticle", mock.Anything, mock.Anything).Return(nil).Once()
		mockRedis.On("ClearAll", mock.Anything).Return(nil).Once()

		usecase := NewArticleUsecase(mockMYSQL, mockRedis)
		err := usecase.PostArticle(context.Background(), subaru)
		assert.Equal(t, err, nil)
		assert.NoError(t, err)
		mockMYSQL.AssertExpectations(t)
	})

	t.Run("Fail - An Error Occured on database", func(t *testing.T) {
		mockMYSQL.On("PostArticle", mock.Anything, mock.Anything).Return(errors.New("query error")).Once()

		usecase := NewArticleUsecase(mockMYSQL, mockRedis)
		err := usecase.PostArticle(context.Background(), subaru)
		assert.Equal(t, err.Error(), "query error")
		assert.Error(t, err)
		mockMYSQL.AssertExpectations(t)
	})
}

func TestGetArticle(t *testing.T) {
	mockMYSQL := new(mocks.ArticleRepository)
	mockRedis := new(mocks.ArticleRedisRepository)
	novel := []domain.Article{
		{Author: "Nagatsuki", Title: "Re-Zero", Body: "An Adventure"},
		{Author: "Kinugasa", Title: "Youjitsu", Body: "A Romance"},
		{Author: "Izushiro", Title: "Saikyou Mahoushi", Body: "A Magic Adventure"},
	}

	t.Run("Success Get data from database", func(t *testing.T) {
		mockRedis.On("GetArticles", mock.Anything, mock.Anything).Return(nil, errors.New("redis: nil")).Once()
		mockMYSQL.On("GetArticles", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(novel, nil).Once()
		mockRedis.On("PostArticleToRedis", mock.Anything, mock.Anything).Return(nil).Once()
		mockRedis.On("PostAllToRedis", mock.Anything, mock.Anything).Return(nil).Once()

		usecase := NewArticleUsecase(mockMYSQL, mockRedis)
		articles, err := usecase.GetArticles(context.Background(), "", "", "")
		assert.Equal(t, articles[2].Title, "Saikyou Mahoushi")
		assert.NotEqual(t, articles[0].Author, "Izushiro")
		assert.NoError(t, err)
		mockMYSQL.AssertExpectations(t)
	})

	t.Run("Success Get data from Redis", func(t *testing.T) {
		mockRedis.On("GetArticles", mock.Anything, mock.Anything).Return(novel, nil).Once()

		usecase := NewArticleUsecase(mockMYSQL, mockRedis)
		articles, err := usecase.GetArticles(context.Background(), "", "", "")
		assert.Equal(t, articles[2].Title, "Saikyou Mahoushi")
		assert.NotEqual(t, articles[0].Author, "Izushiro")
		assert.NoError(t, err)
		mockMYSQL.AssertExpectations(t)
	})

	t.Run("An Error Occured on Redis", func(t *testing.T) {
		mockRedis.On("GetArticles", mock.Anything, mock.Anything).Return(nil, errors.New("some error")).Once()

		usecase := NewArticleUsecase(mockMYSQL, mockRedis)
		articles, err := usecase.GetArticles(context.Background(), "", "", "")
		assert.Nil(t, articles)
		assert.Error(t, err)
		mockMYSQL.AssertExpectations(t)
	})

	t.Run("An Error Occured when trying to get data from database", func(t *testing.T) {
		mockRedis.On("GetArticles", mock.Anything, mock.Anything).Return(nil, errors.New("redis: nil")).Once()
		mockMYSQL.On("GetArticles", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(novel, nil).Once()
		mockRedis.On("PostArticleToRedis", mock.Anything, mock.Anything).Return(errors.New("error save to redis")).Once()

		usecase := NewArticleUsecase(mockMYSQL, mockRedis)
		articles, err := usecase.GetArticles(context.Background(), "", "", "")
		assert.Nil(t, articles)
		assert.Equal(t, err.Error(), "error save to redis")
		assert.Error(t, err)
		mockMYSQL.AssertExpectations(t)
	})

	t.Run("An Error Occured when trying to post data to redis", func(t *testing.T) {
		mockRedis.On("GetArticles", mock.Anything, mock.Anything).Return(nil, errors.New("redis: nil")).Once()
		mockMYSQL.On("GetArticles", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("database error")).Once()

		usecase := NewArticleUsecase(mockMYSQL, mockRedis)
		articles, err := usecase.GetArticles(context.Background(), "", "", "")
		assert.Nil(t, articles)
		assert.Error(t, err)
		mockMYSQL.AssertExpectations(t)
	})
}
