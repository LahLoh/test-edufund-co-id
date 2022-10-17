package handler

import (
	"encoding/json"
	"net/http"

	"github.com/shandysiswandi/test-edufund-co-id/model"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq model.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		errorResponse(w, err)
		return
	}

	token, err := h.svc.Login(r.Context(), loginReq)
	if err != nil {
		errorResponse(w, err)
		return
	}

	jsonResponse(w, http.StatusOK, model.LoginResponse{
		AccessToken: token,
	})
}
