package service

import (
	"context"
	"errors"
	"io"
	"regexp"
	"time"

	"github.com/shandysiswandi/test-edufund-co-id/model"
)

type Repository interface {
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
	CreateUser(ctx context.Context, user model.User) error
	io.Closer
}

type Hasher interface {
	Hash(str string) (string, error)
	Verify(str, hash string) bool
}

type Token interface {
	Sign(exp time.Time, data any) (string, error)
}

type Clocker interface {
	Now() time.Time
}

type Service interface {
	Register(ctx context.Context, req model.RegisterRequest) error
	Login(ctx context.Context, req model.LoginRequest) (string, error)
}

var (
	ErrFullname                  = errors.New("name should be 2 characters or more")
	ErrUsername                  = errors.New("please provide a valid email address")
	ErrPassword                  = errors.New("password should be at least 12 characters long")
	ErrConfirmationPassword      = errors.New("confirmation password does not match")
	ErrInternal                  = errors.New("internal server error")
	ErrInvalidUsernameOrPassword = errors.New("invalid username / password")
)

type service struct {
	repo       Repository
	hasher     Hasher
	token      Token
	clocker    Clocker
	regexEmail *regexp.Regexp
}

func New(repo Repository, hasher Hasher, token Token, clocker Clocker) *service {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return &service{
		repo:       repo,
		hasher:     hasher,
		token:      token,
		clocker:    clocker,
		regexEmail: re,
	}
}
