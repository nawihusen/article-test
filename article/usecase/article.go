package usecase

import (
	"alpha-test/domain"
	"context"
	"time"
	// log "github.com/sirupsen/logrus"
)

// articleUsecase is struct usecase
type articleUsecase struct {
	articleRepo domain.ArticleRepository
	redisRepo   domain.ArticleRedisRepository
}

// NewAccountUsecase is constructor of account usecase
func NewArticleUsecase(articleRepo domain.ArticleRepository, redisRepo domain.ArticleRedisRepository) domain.ArticleUsecase {
	return &articleUsecase{
		articleRepo: articleRepo,
		redisRepo:   redisRepo,
	}
}

func (au *articleUsecase) PostArticle(ctx context.Context, article domain.Article) error {
	// add article to database
	err := au.articleRepo.PostArticle(ctx, article)

	// clear chache
	err = au.redisRepo.ClearAll(ctx)

	return err
}

func (au *articleUsecase) GetArticles(ctx context.Context, author, title, body string) ([]domain.Article, error) {
	// get from chache and return if article data exist
	articles, err := au.redisRepo.GetArticles(ctx, title)
	if err != nil {
		return nil, err
	}
	if articles != nil {
		return articles, err
	}

	// get from database
	articles, err = au.articleRepo.GetArticles(ctx, author, title, body)
	if err != nil {
		return nil, err
	}

	// add to chache
	err = au.redisRepo.PostArticleToRedis(ctx, articles)

	return articles, err
}

func (au *articleUsecase) Test(ctx context.Context) error {
	data := domain.Article{
		ID:      1,
		Author:  "author",
		Title:   "title",
		Body:    "body",
		Created: time.Now(),
	}
	err := au.redisRepo.Test(ctx, data)
	return err
}
