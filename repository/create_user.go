package repository

import (
	"context"

	"github.com/shandysiswandi/test-edufund-co-id/model"
)

var (
	createUser = `INSERT INTO users(fullname, username, password) VALUES(?,?,?);`
)

func (r *repository) CreateUser(ctx context.Context, user model.User) error {
	_, err := r.db.ExecContext(ctx, createUser, user.Fullname, user.Username, user.Password)
	return err
}
