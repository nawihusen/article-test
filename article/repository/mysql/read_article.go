package mysql

import (
	"alpha-test/domain"
	"context"
	"database/sql"
	"time"

	log "github.com/sirupsen/logrus"
)

type mysqlArticleRepository struct {
	Conn *sql.DB
}

// NewMySQLArticleRepository is constructor of MySQL repository
func NewMySQLArticleRepository(Conn *sql.DB) domain.ArticleRepository {
	return &mysqlArticleRepository{Conn}
}

func (db *mysqlArticleRepository) GetArticles(ctx context.Context, author, title, body string) ([]domain.Article, error) {
	query := `SELECT id, author, title, body, created FROM article`
	var params []interface{}
	if author != "" || title != "" || body != "" {
		query += ` WHERE `
	}

	if author != "" {
		query += ` author = ? `
		params = append(params, author)
	}

	if title != "" {
		if author != "" {
			query += ` AND `
		}

		query += ` title LIKE "%` + title + `%" `
	}

	if body != "" {
		if author != "" || title != "" {
			query += ` AND `
		}

		query += ` body LIKE "%` + body + `%" `
	}

	query += " ORDER BY created ASC"

	log.Debug(query)

	rows, err := db.Conn.QueryContext(ctx, query, params...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	var at []domain.Article
	for rows.Next() {
		var i domain.Article
		var temp time.Time
		err := rows.Scan(&i.ID, &i.Author, &i.Title, &i.Body, &temp)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		i.Created = temp.Format(time.RFC3339)

		at = append(at, i)
	}

	return at, nil
}
