package usecase

import (
	"alpha-test/domain"
	"context"

	log "github.com/sirupsen/logrus"
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
	if err != nil {
		return err
	}

	// clear chache
	err = au.redisRepo.ClearAll(ctx)

	return err
}

func (au *articleUsecase) GetArticles(ctx context.Context, author, title, body string) (articles []domain.Article, err error) {
	// get from chache and return if article data exist
	if author == "" && body == "" {
		articles, err = au.redisRepo.GetArticles(ctx, title)
		if err != nil {
			if err.Error() == "redis: nil" {
				log.Debug("No data on redis, Go to database")
			} else {
				return nil, err
			}
		}

		// return data from redis if data exist
		if articles != nil {
			return articles, err
		}
	}
	// fmt.Println("test on redis <=== this should not print")

	// get from database
	articles, err = au.articleRepo.GetArticles(ctx, author, title, body)
	if err != nil {
		return nil, err
	}

	// add one-one to chache
	err = au.redisRepo.PostArticleToRedis(ctx, articles)
	if err != nil {
		return nil, err
	}

	// add all to chache when have no param
	if author == "" && title == "" && body == "" {
		err = au.redisRepo.PostAllToRedis(ctx, articles)
	}

	return articles, err
}
