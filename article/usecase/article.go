package usecase

import (
	"alpha-test/domain"
	"context"
	// log "github.com/sirupsen/logrus"
)

// articleUsecase is struct usecase
type articleUsecase struct {
	articleRepo domain.ArticleRepository
	redisRepo   domain.ArticleRedisRepository
}

// NewAccountUsecase is constructor of account usecase
func NewArticleUsecase(articleRepo domain.ArticleRepository) domain.ArticleUsecase {
	return &articleUsecase{
		articleRepo: articleRepo,
	}
}

func (au *articleUsecase) PostArticle(ctx context.Context, article domain.Article) error {
	err := au.articleRepo.PostArticle(ctx, article)
	return err
}

func (au *articleUsecase) GetArticles(ctx context.Context, author, title, body string) ([]domain.Article, error) {
	// get from chache and return if article data exist
	articles, err := au.redisRepo.GetArticles(ctx, author, title)
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
