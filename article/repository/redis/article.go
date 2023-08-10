package redis

import (
	"alpha-test/domain"
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

type redisArticleRepository struct {
	Conn *redis.Client
}

// NewRedisArticleRepository is constructor of Redis repository
func NewRedisArticleRepository(Conn *redis.Client) domain.ArticleRedisRepository {
	return &redisArticleRepository{Conn}
}

// PostArticleToRedis is a method to save an article to Redis
func (r *redisArticleRepository) PostArticleToRedis(ctx context.Context, article []domain.Article) (err error) {
	expire := time.Duration(viper.GetInt("redis.expire") * int(time.Hour))

	for _, v := range article {
		pipe := r.Conn.Pipeline()
		pipe.HSet(ctx, v.Title, "id", v.ID, "author", v.Author, "title", v.Title, "body", v.Body, "created", v.Created)

		pipe.Expire(ctx, v.Title, expire)

		_, err = pipe.Exec(ctx)
	}

	return err
}

// PostAllToRedis is a method to save an article to Redis
func (r *redisArticleRepository) PostAllToRedis(ctx context.Context, article []domain.Article) (err error) {
	all, err := json.Marshal(article)
	if err != nil {
		return err
	}
	expire := time.Duration(viper.GetInt("redis.expire") * int(time.Hour))

	_, err = r.Conn.Set(ctx, "all", all, expire).Result()

	return err
}

// GetArticles is a method to get data from Redis
func (r *redisArticleRepository) GetArticles(ctx context.Context, title string) (articles []domain.Article, err error) {
	if title != "" {
		data := r.Conn.HGetAll(ctx, title)
		_, err = data.Result()
		if err != nil {
			return
		}

		temp := domain.Article{}
		err = data.Scan(&temp)
		if err != nil {
			return
		}

		if temp.ID == 0 {
			return nil, errors.New("redis: nil")
		}

		articles = append(articles, temp)
	} else {
		res, err := r.Conn.Get(ctx, "all").Bytes()
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(res, &articles)
	}

	return articles, nil
}

// ClearAll is a method to clear all Redis data
func (r *redisArticleRepository) ClearAll(ctx context.Context) (err error) {
	_, err = r.Conn.Del(ctx, "all").Result()
	return err
}
