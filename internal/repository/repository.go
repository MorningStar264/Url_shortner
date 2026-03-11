package repository

import "github.com/MorningStar264/Url_shortner/internal/server"

type Repositories struct {
	UserMethods *UserMethods
	LinkMethods *LinkMethods
}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{
		UserMethods: NewUserMethods(s),
		LinkMethods: NewLinkMethods(s),
	}
}
