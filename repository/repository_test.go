package repository

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
)

func TestNew(t *testing.T) {
	type args struct {
		opt MysqlOption
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ErrorOpenConnection",
			args: args{
				opt: MysqlOption{
					Username: "root",
					Password: "",
					Host:     "localhost",
					Port:     "3306",
					Database: "test",
					Options:  "charset=utf8mb4&parseTime=True&loc=Asia/Jakarta",
				},
			},
			wantErr: true,
		},
		{
			name: "SuccessOpenConnection",
			args: args{
				opt: MysqlOption{
					Username: "root",
					Password: "",
					Host:     "localhost",
					Port:     "3306",
					Database: "test",
					Options:  "charset=utf8mb4&parseTime=True&loc=UTC",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := New(tt.args.opt); (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_repository_Ping(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.MonitorPingsOption(true))
	defer db.Close()

	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		mockFn  func() fields
		wantErr bool
	}{
		{
			name: "ErrorPing",
			mockFn: func() fields {
				mock.ExpectPing().WillReturnError(errors.New("fake"))

				return fields{
					db: db,
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: tt.mockFn().db,
			}
			if err := r.Ping(); (err != nil) != tt.wantErr {
				t.Errorf("repository.Ping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repository_Close(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		wantErr bool
		mockFn  func() fields
	}{
		{
			name: "ErrorClose",
			mockFn: func() fields {
				mock.ExpectClose().WillReturnError(errors.New("fake"))

				return fields{
					db: db,
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: tt.mockFn().db,
			}
			if err := r.Close(); (err != nil) != tt.wantErr {
				t.Errorf("repository.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repository_Migrate(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ErrorMigrate",
			fields: fields{
				db: &sql.DB{},
			},
			args: args{
				dir: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: tt.fields.db,
			}
			if err := r.Migrate(tt.args.dir); (err != nil) != tt.wantErr {
				t.Errorf("repository.Migrate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
