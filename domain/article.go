package domain

import (
	"context"
	"time"
)

type Article struct {
	ID      int64     `json:"id" redis:"id"`
	Author  string    `json:"author" redis:"author" validate:"required"`
	Title   string    `json:"title" redis:"title" validate:"required"`
	Body    string    `json:"body" redis:"body" validate:"required"`
	Created time.Time `json:"created" redis:"created"`
}

// Usecase is Article usecase
type ArticleUsecase interface {
	PostArticle(ctx context.Context, article Article) (err error)
	GetArticles(ctx context.Context, author, title, body string) (articles []Article, err error)
}

type ArticleRepository interface {
	PostArticle(ctx context.Context, article Article) (err error)
	GetArticles(ctx context.Context, author, title, body string) (articles []Article, err error)
}

type ArticleRedisRepository interface {
	PostArticleToRedis(ctx context.Context, article []Article) (err error)
	GetArticles(ctx context.Context, author, title string) (articles []Article, err error)
}
