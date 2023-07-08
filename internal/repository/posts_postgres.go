package repository

import (
	"github.com/jackc/pgx"
	"twitter-go/entity"
)

type PostRepository struct {
	db *pgx.Conn
}

func NewPostPostgres(db *pgx.Conn) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(data entity.CreatePostRecord) (int, error) {
	var postId int
	query := "INSERT INTO posts (text, created_by) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRow(query, data.Text, data.CreatedBy).Scan(&postId)

	if err != nil {
		return 0, err
	}

	return postId, nil
}

func (r *PostRepository) FindOne(id int) (*entity.Post, error) {
	var post *entity.Post
	query := "SELECT id, text, created_by, created_at FROM posts WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&post.Id, &post.Text, &post.CreatedBy, &post.CreatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return post, nil
}

func (r *PostRepository) FindByUserId(userId int) ([]entity.Post, error) {
	query := "SELECT id, text, created_at, created_by FROM posts AS p WHERE p.created_by = $1"
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	posts := make([]entity.Post, 0)

	for rows.Next() {
		var post entity.Post
		err := rows.Scan(&post.Id, &post.Text, &post.CreatedAt, &post.CreatedBy)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
