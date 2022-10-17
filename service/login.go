package service

import (
	"context"
	"net/http"
	"time"

	"github.com/shandysiswandi/test-edufund-co-id/model"
)

// Login is usecase for user login
// for validation i dont use validation library
func (s *service) Login(ctx context.Context, req model.LoginRequest) (string, error) {
	if ok := s.regexEmail.MatchString(req.Username); !ok {
		return "", model.ErrorResponse{
			Msg:  ErrUsername.Error(),
			Code: http.StatusBadRequest,
		}
	}

	if len(req.Password) < 12 {
		return "", model.ErrorResponse{
			Msg:  ErrPassword.Error(),
			Code: http.StatusBadRequest,
		}
	}

	user, err := s.repo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return "", model.ErrorResponse{
			Msg:  ErrInvalidUsernameOrPassword.Error(),
			Code: http.StatusUnauthorized,
		}
	}

	if ok := s.hasher.Verify(req.Password, user.Password); !ok {
		return "", model.ErrorResponse{
			Msg:  ErrInvalidUsernameOrPassword.Error(),
			Code: http.StatusUnauthorized,
		}
	}

	accToken, err := s.token.Sign(s.clocker.Now().Add(10*time.Minute), user)
	if err != nil {
		return "", model.ErrorResponse{
			Msg:  ErrInternal.Error(),
			Code: http.StatusInternalServerError,
		}
	}

	return accToken, nil
}
