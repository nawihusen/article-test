package mysql

import (
	"alpha-test/domain"
	"context"
)

func (db *mysqlArticleRepository) PostArticle(ctx context.Context, article domain.Article) error {
	query := `INSERT INTO article (author, title, body, created) VALUES (?, ?, ?, NOW())`

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, article.Author, article.Title, article.Body)
	if err != nil {
		return err
	}

	return nil
}
