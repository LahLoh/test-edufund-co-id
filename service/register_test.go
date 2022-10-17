package service

import (
	"context"
	"errors"
	"net/http"
	"regexp"
	"testing"

	"github.com/shandysiswandi/test-edufund-co-id/mock"
	"github.com/shandysiswandi/test-edufund-co-id/model"
	"github.com/stretchr/testify/assert"
)

func Test_service_Register(t *testing.T) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	type fields struct {
		repo    Repository
		hasher  Hasher
		token   Token
		clocker Clocker
	}

	type args struct {
		ctx context.Context
		req model.RegisterRequest
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		mockFn  func(a args) fields
	}{
		{
			name: "ErrorFullname",
			args: args{
				ctx: context.Background(),
				req: model.RegisterRequest{
					Fullname:             "",
					Username:             "",
					Password:             "",
					ConfirmationPassword: "",
				},
			},
			wantErr: model.ErrorResponse{
				Msg:  ErrFullname.Error(),
				Code: http.StatusBadRequest,
			},
			mockFn: func(args) fields {
				return fields{}
			},
		},
		{
			name: "ErrorEmailOrUsername",
			args: args{
				ctx: context.Background(),
				req: model.RegisterRequest{
					Fullname:             "Admin",
					Username:             "",
					Password:             "",
					ConfirmationPassword: "",
				},
			},
			wantErr: model.ErrorResponse{
				Msg:  ErrUsername.Error(),
				Code: http.StatusBadRequest,
			},
			mockFn: func(args) fields {
				return fields{}
			},
		},
		{
			name: "ErrorPassword",
			args: args{
				ctx: context.Background(),
				req: model.RegisterRequest{
					Fullname:             "Admin",
					Username:             "admi@admin.com",
					Password:             "",
					ConfirmationPassword: "",
				},
			},
			wantErr: model.ErrorResponse{
				Msg:  ErrPassword.Error(),
				Code: http.StatusBadRequest,
			},
			mockFn: func(args) fields {
				return fields{}
			},
		},
		{
			name: "ErrorConfirmationPassword",
			args: args{
				ctx: context.Background(),
				req: model.RegisterRequest{
					Fullname:             "Admin",
					Username:             "admi@admin.com",
					Password:             "password_password",
					ConfirmationPassword: "",
				},
			},
			wantErr: model.ErrorResponse{
				Msg:  ErrConfirmationPassword.Error(),
				Code: http.StatusBadRequest,
			},
			mockFn: func(args) fields {
				return fields{}
			},
		},
		{
			name: "ErrorHasherHash",
			args: args{
				ctx: context.Background(),
				req: model.RegisterRequest{
					Fullname:             "Admin",
					Username:             "admi@admin.com",
					Password:             "password_password",
					ConfirmationPassword: "password_password",
				},
			},
			wantErr: model.ErrorResponse{
				Msg:  ErrInternal.Error(),
				Code: http.StatusBadRequest,
			},
			mockFn: func(a args) fields {
				mockHasher := mock.NewHasher(t)

				mockHasher.On("Hash", a.req.Password).Return("", errors.New("fake"))

				return fields{
					hasher: mockHasher,
				}
			},
		},
		{
			name: "ErrorRepoCreateUser",
			args: args{
				ctx: context.Background(),
				req: model.RegisterRequest{
					Fullname:             "Admin",
					Username:             "admi@admin.com",
					Password:             "password_password",
					ConfirmationPassword: "password_password",
				},
			},
			wantErr: model.ErrorResponse{
				Msg:  ErrInternal.Error(),
				Code: http.StatusBadRequest,
			},
			mockFn: func(a args) fields {
				mockHasher := mock.NewHasher(t)
				mockRepo := mock.NewRepository(t)

				pass := "hash(password)"
				mockHasher.On("Hash", a.req.Password).Return(pass, nil)

				user := model.User{
					Fullname: a.req.Fullname,
					Username: a.req.Username,
					Password: pass,
				}
				mockRepo.On("CreateUser", a.ctx, user).Return(errors.New("fake"))

				return fields{
					repo:   mockRepo,
					hasher: mockHasher,
				}
			},
		},
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				req: model.RegisterRequest{
					Fullname:             "Admin",
					Username:             "admi@admin.com",
					Password:             "password_password",
					ConfirmationPassword: "password_password",
				},
			},
			wantErr: nil,
			mockFn: func(a args) fields {
				mockHasher := mock.NewHasher(t)
				mockRepo := mock.NewRepository(t)

				pass := "hash(password)"
				mockHasher.On("Hash", a.req.Password).Return(pass, nil)

				user := model.User{
					Fullname: a.req.Fullname,
					Username: a.req.Username,
					Password: pass,
				}
				mockRepo.On("CreateUser", a.ctx, user).Return(nil)

				return fields{
					repo:   mockRepo,
					hasher: mockHasher,
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

			err := s.Register(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}
