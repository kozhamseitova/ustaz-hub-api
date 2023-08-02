package handler

import "github.com/kozhamseitova/ustaz-hub-api/internal/service"

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}
