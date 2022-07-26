package handler

import (
	"github.com/Vallghall/mt/test/internal/service"
)

type Handler struct {
	s *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{s}
}
