package db

import (
	"context"
	"database/sql"
	"log"

	"github.com/JalMurH/DockerDevDeploy/rest-ws/models"
	_ "github.com/lib/pq"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(url string) (*PostgresRepo, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepo{db}, nil
}

func (repo *PostgresRepo) InsertUser(ctx context.Context, user *models.User) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (id, email, password) VALUES($1,$2, $3)", user.Id, user.Email, user.Password)
	return err
}

func (repo *PostgresRepo) InsertPost(ctx context.Context, post *models.Post) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO posts (id, post_content, user_id) VALUES($1,$2, $3)", post.Id, post.PostContent, post.UserId)
	return err
}

func (repo *PostgresRepo) GetUserById(ctx context.Context, id string) (*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, email FROM users WHERE id=$1", id)

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var user = models.User{}

	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email); err == nil {
			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *PostgresRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, email, password FROM users WHERE email=$1", email)

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var user = models.User{}

	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email, &user.Password); err == nil {
			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *PostgresRepo) Close() error {
	return repo.db.Close()
}
