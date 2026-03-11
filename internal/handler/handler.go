package handler

import (
	"github.com/MorningStar264/Url_shortner/internal/repository"
	"github.com/MorningStar264/Url_shortner/internal/server"
)

type Handlers struct {
	UserHandler *UserHandler
	LinkHandler *LinkHandler
}

func NewHandlers(srv *server.Server, repo *repository.Repositories) *Handlers {
	return &Handlers{
		UserHandler: NewUserHandler(srv, repo),
		LinkHandler: NewLinkHandler(srv, repo),
	}
}
