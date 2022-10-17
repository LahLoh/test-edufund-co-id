package handler

import "github.com/shandysiswandi/test-edufund-co-id/service"

type Handler struct {
	svc service.Service
}

func New(svc service.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}
