package service

import (
	"context"
	"net/http"

	"github.com/shandysiswandi/test-edufund-co-id/model"
)

// Register is usecase for user registration
// for validation i dont use validation library
func (s *service) Register(ctx context.Context, req model.RegisterRequest) error {
	if len(req.Fullname) < 2 {
		return model.ErrorResponse{
			Msg:  ErrFullname.Error(),
			Code: http.StatusBadRequest,
		}
	}

	if ok := s.regexEmail.MatchString(req.Username); !ok {
		return model.ErrorResponse{
			Msg:  ErrUsername.Error(),
			Code: http.StatusBadRequest,
		}
	}

	if len(req.Password) < 12 {
		return model.ErrorResponse{
			Msg:  ErrPassword.Error(),
			Code: http.StatusBadRequest,
		}
	}

	if req.Password != req.ConfirmationPassword {
		return model.ErrorResponse{
			Msg:  ErrConfirmationPassword.Error(),
			Code: http.StatusBadRequest,
		}
	}

	pass, err := s.hasher.Hash(req.Password)
	if err != nil {
		return model.ErrorResponse{
			Msg:  ErrInternal.Error(),
			Code: http.StatusBadRequest,
		}
	}

	if err = s.repo.CreateUser(ctx, model.User{
		Fullname: req.Fullname,
		Username: req.Username,
		Password: pass,
	}); err != nil {
		return model.ErrorResponse{
			Msg:  ErrInternal.Error(),
			Code: http.StatusBadRequest,
		}
	}

	return nil
}
