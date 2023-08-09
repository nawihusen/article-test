package redis

import (
	"alpha-test/domain"
	"context"

	"github.com/go-redis/redis/v8"
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
	// var params []interface{}
	// params = append(params, "recruiter_id", add.RecruiterID, "saksi_id", add.SaksiID, "coordinator_id", add.CoordinatorID, "tps_no", add.TpsNo, "attendance", add.Attendance, "recruiter_role", add.RoleRecruiter, "saksi_role", add.RoleSaksi, "coordinator_role", add.RoleCoordinator)
	// _, err = r.Conn.HSet(ctx, token, params...).Result()

	// _, err = r.Conn.Set(ctx, token, role, 0).Result()
	return
}

// GetArticles is a method to get data from Redis
func (r *redisArticleRepository) GetArticles(ctx context.Context, author, title string) (articles []domain.Article, err error) {
	// var params []interface{}
	// params = append(params, "recruiter_id", add.RecruiterID, "saksi_id", add.SaksiID, "coordinator_id", add.CoordinatorID, "tps_no", add.TpsNo, "attendance", add.Attendance, "recruiter_role", add.RoleRecruiter, "saksi_role", add.RoleSaksi, "coordinator_role", add.RoleCoordinator)
	// _, err = r.Conn.HSet(ctx, token, params...).Result()

	// _, err = r.Conn.Set(ctx, token, role, 0).Result()
	return
}

// ClearAuthorArticle is a method to clear data whenever an author make a new article
func (r *redisArticleRepository) ClearAuthorArticle(ctx context.Context, author string) (err error) {
	// var params []interface{}
	// params = append(params, "recruiter_id", add.RecruiterID, "saksi_id", add.SaksiID, "coordinator_id", add.CoordinatorID, "tps_no", add.TpsNo, "attendance", add.Attendance, "recruiter_role", add.RoleRecruiter, "saksi_role", add.RoleSaksi, "coordinator_role", add.RoleCoordinator)
	// _, err = r.Conn.HSet(ctx, token, params...).Result()

	// _, err = r.Conn.Set(ctx, token, role, 0).Result()
	return
}

// ClearAll is a method to clear all Redis data
func (r *redisArticleRepository) ClearAll(ctx context.Context) (err error) {
	// var params []interface{}
	// params = append(params, "recruiter_id", add.RecruiterID, "saksi_id", add.SaksiID, "coordinator_id", add.CoordinatorID, "tps_no", add.TpsNo, "attendance", add.Attendance, "recruiter_role", add.RoleRecruiter, "saksi_role", add.RoleSaksi, "coordinator_role", add.RoleCoordinator)
	// _, err = r.Conn.HSet(ctx, token, params...).Result()

	// _, err = r.Conn.Set(ctx, token, role, 0).Result()
	return
}

// // SetSession is save to Redis if any user login
// func (r *redisAccountRepository) SetSession(ctx context.Context, token string, add domain.SaksiAdditionalData) (err error) {
// 	var params []interface{}
// 	params = append(params, "recruiter_id", add.RecruiterID, "saksi_id", add.SaksiID, "coordinator_id", add.CoordinatorID, "tps_no", add.TpsNo, "attendance", add.Attendance, "recruiter_role", add.RoleRecruiter, "saksi_role", add.RoleSaksi, "coordinator_role", add.RoleCoordinator)
// 	_, err = r.Conn.HSet(ctx, token, params...).Result()

// 	// _, err = r.Conn.Set(ctx, token, role, 0).Result()
// 	return
// }

// // GetSession is get role from Redis if still have session
// func (r *redisAccountRepository) GetSession(ctx context.Context, token string) (add domain.SaksiAdditionalData, err error) {
// 	data := r.Conn.HGetAll(ctx, token)
// 	res, err := data.Result()
// 	if err != nil {
// 		return
// 	}

// 	if len(res) == 0 {
// 		err = errors.New("not found")
// 		return
// 	}

// 	err = data.Scan(&add)
// 	if err != nil {
// 		return
// 	}

// 	// role, err = r.Conn.Get(ctx, token).Result()
// 	return
// }

// // DeleteSession id delete session (key) if have session no more from auth
// func (r *redisAccountRepository) DeleteSession(ctx context.Context, token string) (err error) {
// 	_, err = r.Conn.Del(ctx, token).Result()

// 	return
// }

// // DeleteSession id delete session (key) if have session no more from auth
// func (r *redisAccountRepository) AttendanceSession(ctx context.Context, token string) (err error) {
// 	_, err = r.Conn.HSet(ctx, token, "attendance", true).Result()

// 	return
// }
