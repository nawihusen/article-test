package domain

import (
	"context"
)

type Article struct {
	ID      int64  `json:"id" form:"id" redis:"id"`
	Author  string `json:"author" form:"author" redis:"author" validate:"required"`
	Title   string `json:"title" form:"title" redis:"title" validate:"required"`
	Body    string `json:"body" form:"body" redis:"body" validate:"required"`
	Created string `json:"created" form:"created" redis:"created"`
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
	GetArticles(ctx context.Context, title string) (articles []Article, err error)
	ClearAll(ctx context.Context) (err error)
	PostAllToRedis(ctx context.Context, article []Article) (err error)
}
