package store

import (
	"context"
	"database/sql"
)

type Comment struct {
	ID int64 `json:"id"`
	PostID int64 `json:"post_id"`
	UserID int64 `json:"user_id"`
	Content string `json:"content"`
	CreatedAt string `json:"created_at"`
	User User `json:"user"`
}

type CommentStore struct {
	db *sql.DB
}

func (s *CommentStore) CreateByPostID(ctx context.Context, comment *Comment) error {
	query := `
		WITH inserted_comment AS (
			INSERT INTO comments (post_id, user_id, content)
			VALUES ($1, $2, $3)
			RETURNING id, post_id, user_id, content, created_at
		)
		SELECT 
			ic.id, ic.post_id, ic.user_id, ic.content, ic.created_at, 
			u.id, u.username
		FROM inserted_comment ic
		JOIN users u ON ic.user_id = u.id
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		comment.PostID,
		comment.UserID,
		comment.Content, 
	).Scan(
		&comment.ID,
		&comment.PostID,
		&comment.UserID,
		&comment.Content,
		&comment.CreatedAt,
		&comment.User.ID,
		&comment.User.Username,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *CommentStore) GetByPostID(ctx context.Context, postID int64) ([]*Comment, error) {
	query := `
		SELECT c.id, c.post_id, c.user_id, c.content, c.created_at, users.username, users.id FROM comments c
		JOIN users on users.id = c.user_id
		WHERE c.post_id = $1
		ORDER BY c.created_at DESC
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(
		ctx,
		query,
		postID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []*Comment{}

	for rows.Next() {
		c := Comment{}
		c.User = User{}
		err := rows.Scan(
			&c.ID,
			&c.PostID,
			&c.UserID,
			&c.Content,
			&c.CreatedAt,
			&c.User.Username,
			&c.User.ID,
		)

		if err != nil {
			return nil, err
		}

		comments = append(comments, &c)
	}

	return comments, nil
}