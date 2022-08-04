package postgres

import (
	"context"
	"database/sql"
	"fmt"

	pb "crud_service/genproto/crud_service"

	"github.com/jmoiron/sqlx"
)

type crudRepo struct {
	db *sqlx.DB
}

func NewCrudRepo(db *sqlx.DB) CrudI {
	return &crudRepo{
		db: db,
	}
}

func (r *crudRepo) GetAll(ctx context.Context, req *pb.ListReq) (*pb.ListRespPost, error) {
	var resp pb.ListRespPost

	rows, err := r.db.Query(`
		SELECT 
			id,
			user_id,
			title,
			body
		FROM posts 
		WHERE deleted_at=0 
		LIMIT $1 
		OFFSET $2`,
		req.Limit,
		req.Page,
	)
	if err != nil {
		return nil, fmt.Errorf("QUERY GET ALL: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var post pb.Post
		err = rows.Scan(
			&post.Id,
			&post.UserId,
			&post.Title,
			&post.Body,
		)
		if err != nil {
			return nil, fmt.Errorf("getall post scan: %w", err)
		}
		resp.Posts = append(resp.Posts, &post)
	}

	err = r.db.QueryRow(`SELECT COUNT(1) FROM posts WHERE deleted_at=0`).Scan(&resp.Total)
	if err != nil {
		return nil, fmt.Errorf("getall count: %w", err)
	}
	return &resp, nil
}

func (r *crudRepo) Get(ctx context.Context, req *pb.ById) (*pb.Post, error) {
	var post pb.Post

	err := r.db.QueryRow(`
	SELECT
		id,
		user_id,
		title,
		body
	FROM posts
	WHERE id = $1 AND deleted_at = 0`,
		req.Id).Scan(
		&post.Id,
		&post.UserId,
		&post.Title,
		&post.Body,
	)
	if err != nil {
		return nil, fmt.Errorf("get attribute: %w", err)
	}

	return &post, nil
}

func (r *crudRepo) Update(ctx context.Context, req *pb.Post) error {
	res, err := r.db.Exec(`
	UPDATE posts SET
		user_id = $2,
		title = $3,
		body = $4,
		updated_at = CURRENT_TIMESTAMP
	WHERE id = $1 AND deleted_at = 0`,
		req.Id,
		req.UserId,
		req.Title,
		req.Body,
	)
	if err != nil {
		return fmt.Errorf("update post: %w", err)
	}

	i, err := res.RowsAffected()
	if i == 0 {
		return sql.ErrNoRows
	}
	if err != nil {
		fmt.Println("RowsAffected Error", err)
	}

	return nil
}

func (r *crudRepo) Delete(ctx context.Context, req *pb.ById) error {
	result, err := r.db.Exec(`
	UPDATE posts SET
		deleted_at = date_part('epoch', CURRENT_TIMESTAMP)::INT
	WHERE id = $1 AND deleted_at = 0`, req.Id)
	if err != nil {
		return err
	}

	i, err := result.RowsAffected()
	if i == 0 {
		return sql.ErrNoRows
	}
	if err != nil {
		fmt.Println("RowsAffected Error", err)
	}
	return nil
}
