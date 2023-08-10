// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	domain "alpha-test/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ArticleRedisRepository is an autogenerated mock type for the ArticleRedisRepository type
type ArticleRedisRepository struct {
	mock.Mock
}

// ClearAll provides a mock function with given fields: ctx
func (_m *ArticleRedisRepository) ClearAll(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetArticles provides a mock function with given fields: ctx, title
func (_m *ArticleRedisRepository) GetArticles(ctx context.Context, title string) ([]domain.Article, error) {
	ret := _m.Called(ctx, title)

	var r0 []domain.Article
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]domain.Article, error)); ok {
		return rf(ctx, title)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []domain.Article); ok {
		r0 = rf(ctx, title)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Article)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostAllToRedis provides a mock function with given fields: ctx, article
func (_m *ArticleRedisRepository) PostAllToRedis(ctx context.Context, article []domain.Article) error {
	ret := _m.Called(ctx, article)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []domain.Article) error); ok {
		r0 = rf(ctx, article)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PostArticleToRedis provides a mock function with given fields: ctx, article
func (_m *ArticleRedisRepository) PostArticleToRedis(ctx context.Context, article []domain.Article) error {
	ret := _m.Called(ctx, article)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []domain.Article) error); ok {
		r0 = rf(ctx, article)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewArticleRedisRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewArticleRedisRepository creates a new instance of ArticleRedisRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewArticleRedisRepository(t mockConstructorTestingTNewArticleRedisRepository) *ArticleRedisRepository {
	mock := &ArticleRedisRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}