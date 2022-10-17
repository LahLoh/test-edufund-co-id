package repository

import (
	"context"
	"database/sql"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/shandysiswandi/test-edufund-co-id/model"
)

func Test_repository_GetUserByUsername(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx      context.Context
		username string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.User
		wantErr bool
		mockFn  func(args)
	}{
		{
			name: "ErrorGetUserByUsername",
			fields: fields{
				db: db,
			},
			args: args{
				ctx:      context.Background(),
				username: "username",
			},
			want: model.User{
				ID:        0,
				Fullname:  "",
				Username:  "",
				Password:  "",
				CreatedAt: time.Time{},
			},
			wantErr: true,
			mockFn: func(a args) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow("id")
				mock.ExpectQuery(regexp.QuoteMeta(getUserByUsername)).WithArgs(a.username).WillReturnRows(rows)
			},
		},
		{
			name: "SuccessGetUserByUsername",
			fields: fields{
				db: db,
			},
			args: args{
				ctx:      context.Background(),
				username: "username",
			},
			want: model.User{
				ID:        1,
				Fullname:  "fullname",
				Username:  "username",
				Password:  "hash(password)",
				CreatedAt: time.Time{},
			},
			wantErr: false,
			mockFn: func(a args) {
				rows := sqlmock.NewRows([]string{"id", "fullname", "username", "password", "created_at"}).
					AddRow(1, "fullname", "username", "hash(password)", time.Time{})
				mock.ExpectQuery(regexp.QuoteMeta(getUserByUsername)).WithArgs(a.username).WillReturnRows(rows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			r := &repository{
				db: tt.fields.db,
			}

			got, err := r.GetUserByUsername(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetUserByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetUserByUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}
