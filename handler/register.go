package handler

import (
	"encoding/json"
	"net/http"

	"github.com/shandysiswandi/test-edufund-co-id/model"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var registerReq model.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&registerReq); err != nil {
		errorResponse(w, err)
		return
	}

	err := h.svc.Register(r.Context(), registerReq)
	if err != nil {
		errorResponse(w, err)
		return
	}

	jsonResponse(w, http.StatusOK, model.RegisterResponse{
		Message: http.StatusText(http.StatusOK),
	})
}
