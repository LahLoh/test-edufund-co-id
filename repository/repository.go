package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
)

type repository struct {
	db *sql.DB
}

type MysqlOption struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
	Options  string // charset=utf8mb4&parseTime=True&loc=Asia/Jakarta
}

func New(opt MysqlOption) (*repository, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		opt.Username,
		opt.Password,
		opt.Host,
		opt.Port,
		opt.Database,
		opt.Options,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return &repository{
		db: db,
	}, nil
}

func (r *repository) Ping() error {
	return r.db.Ping()
}

func (r *repository) Migrate(dir string) error {
	goose.SetDialect("mysql")
	return goose.Up(r.db, dir)
}

func (r *repository) Close() error {
	return r.db.Close()
}
