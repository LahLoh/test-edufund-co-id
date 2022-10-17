package repository

import (
	"context"

	"github.com/shandysiswandi/test-edufund-co-id/model"
)

var (
	getUserByUsername = `SELECT id, fullname, username, password, created_at FROM users WHERE username = ? LIMIT 1;`
)

func (r *repository) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	user := model.User{}

	err := r.db.QueryRowContext(ctx, getUserByUsername, username).
		Scan(&user.ID, &user.Fullname, &user.Username, &user.Password, &user.CreatedAt)

	return user, err
}
