package service

import (
	"context"
	"errors"
	"net/http"
	"regexp"
	"testing"
	"time"

	"github.com/shandysiswandi/test-edufund-co-id/mock"
	"github.com/shandysiswandi/test-edufund-co-id/model"
	"github.com/stretchr/testify/assert"
)

func Test_service_Login(t *testing.T) {

	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	type fields struct {
		repo    Repository
		hasher  Hasher
		token   Token
		clocker Clocker
	}

	type args struct {
		ctx context.Context
		req model.LoginRequest
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
		mockFn  func(args) fields
	}{
		{
			name: "ErrorInvalidEmail",
			args: args{
				ctx: context.Background(),
				req: model.LoginRequest{
					Username: "invalid",
					Password: "password_",
				},
			},
			want: "",
			wantErr: model.ErrorResponse{
				Msg:  ErrUsername.Error(),
				Code: http.StatusBadRequest,
			},
			mockFn: func(args) fields {
				return fields{}
			},
		},
		{
			name: "ErrorInvalidPassword",
			args: args{
				ctx: context.Background(),
				req: model.LoginRequest{
					Username: "email@email.com",
					Password: "password_",
				},
			},
			want: "",
			wantErr: model.ErrorResponse{
				Msg:  ErrPassword.Error(),
				Code: http.StatusBadRequest,
			},
			mockFn: func(args) fields {
				return fields{}
			},
		},
		{
			name: "ErrorRepoGetUserByUsername",
			args: args{
				ctx: context.Background(),
				req: model.LoginRequest{
					Username: "email@email.com",
					Password: "password_password",
				},
			},
			want: "",
			wantErr: model.ErrorResponse{
				Msg:  ErrInvalidUsernameOrPassword.Error(),
				Code: http.StatusUnauthorized,
			},
			mockFn: func(a args) fields {
				mockRepo := mock.NewRepository(t)

				mockRepo.On("GetUserByUsername", a.ctx, a.req.Username).Return(model.User{}, errors.New("fake"))

				return fields{
					repo: mockRepo,
				}
			},
		},
		{
			name: "ErrorHasherVerify",
			args: args{
				ctx: context.Background(),
				req: model.LoginRequest{
					Username: "email@email.com",
					Password: "password_password",
				},
			},
			want: "",
			wantErr: model.ErrorResponse{
				Msg:  ErrInvalidUsernameOrPassword.Error(),
				Code: http.StatusUnauthorized,
			},
			mockFn: func(a args) fields {
				mockRepo := mock.NewRepository(t)
				hasher := mock.NewHasher(t)
				user := model.User{}

				mockRepo.On("GetUserByUsername", a.ctx, a.req.Username).Return(user, nil)

				hasher.On("Verify", a.req.Password, user.Password).Return(false)

				return fields{
					repo:   mockRepo,
					hasher: hasher,
				}
			},
		},
		{
			name: "ErrorTokenSign",
			args: args{
				ctx: context.Background(),
				req: model.LoginRequest{
					Username: "email@email.com",
					Password: "password_password",
				},
			},
			want: "",
			wantErr: model.ErrorResponse{
				Msg:  ErrInternal.Error(),
				Code: http.StatusInternalServerError,
			},
			mockFn: func(a args) fields {
				mockRepo := mock.NewRepository(t)
				mockHasher := mock.NewHasher(t)
				mockToken := mock.NewToken(t)
				mockClock := mock.NewClocker(t)
				user := model.User{}

				mockRepo.On("GetUserByUsername", a.ctx, a.req.Username).Return(user, nil)

				mockHasher.On("Verify", a.req.Password, user.Password).Return(true)

				now := time.Now()
				mockClock.On("Now").Return(now)

				mockToken.On("Sign", now.Add(10*time.Minute), user).Return("", errors.New("fake"))

				return fields{
					repo:    mockRepo,
					hasher:  mockHasher,
					token:   mockToken,
					clocker: mockClock,
				}
			},
		},
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				req: model.LoginRequest{
					Username: "email@email.com",
					Password: "password_password",
				},
			},
			want:    "JWT",
			wantErr: nil,
			mockFn: func(a args) fields {
				mockRepo := mock.NewRepository(t)
				mockHasher := mock.NewHasher(t)
				mockToken := mock.NewToken(t)
				mockClock := mock.NewClocker(t)
				user := model.User{}

				mockRepo.On("GetUserByUsername", a.ctx, a.req.Username).Return(user, nil)

				mockHasher.On("Verify", a.req.Password, user.Password).Return(true)

				now := time.Now()
				mockClock.On("Now").Return(now)

				mockToken.On("Sign", now.Add(10*time.Minute), user).Return("JWT", nil)

				return fields{
					repo:    mockRepo,
					hasher:  mockHasher,
					token:   mockToken,
					clocker: mockClock,
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fields := tt.mockFn(tt.args)

			s := &service{
				repo:       fields.repo,
				regexEmail: re,
				hasher:     fields.hasher,
				token:      fields.token,
				clocker:    fields.clocker,
			}

			accToken, err := s.Login(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, accToken)
		})
	}
}
