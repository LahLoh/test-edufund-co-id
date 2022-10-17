package repository

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/shandysiswandi/test-edufund-co-id/model"
)

func Test_repository_CreateUser(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx  context.Context
		user model.User
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mockFn  func(args)
	}{
		{
			name: "ErrorCreateUser",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.Background(),
				user: model.User{
					ID:        0,
					Fullname:  "",
					Username:  "",
					Password:  "",
					CreatedAt: time.Time{},
				},
			},
			wantErr: true,
			mockFn: func(args) {
				mock.ExpectExec(regexp.QuoteMeta(createUser)).WillReturnError(sql.ErrConnDone)
			},
		},
		{
			name: "SuccessCreateUser",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.Background(),
				user: model.User{
					ID:        1,
					Fullname:  "fullname",
					Username:  "username",
					Password:  "password",
					CreatedAt: time.Time{},
				},
			},
			wantErr: false,
			mockFn: func(args) {
				mock.ExpectExec(regexp.QuoteMeta(createUser)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			r := &repository{
				db: tt.fields.db,
			}

			if err := r.CreateUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("repository.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
