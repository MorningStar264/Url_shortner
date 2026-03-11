package router

import (
	"github.com/MorningStar264/Url_shortner/internal/handler"
	// "github.com/MorningStar264/Url_shortner/internal/middlewares"
	"github.com/MorningStar264/Url_shortner/internal/server"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(s *server.Server, h *handler.Handlers) *chi.Mux {

	r := chi.NewRouter()

	//add middleware here
	r.Use(middleware.Logger)

	r.Route("/", func(r chi.Router) {

		//Routes for user mangement
		r.Route("/user", func(r chi.Router) {

			//dummy user username Admin,password Admin123, email admin@gmail.com
			r.Post("/", h.UserHandler.SignUp)
			r.Get("/", h.UserHandler.Login)
			r.Patch("/{userID}", h.UserHandler.UpdateUser)
			r.Delete("/{userId}", h.UserHandler.DeleteUser)
		})

		//Route for shortening
		r.Route("/shorten", func(r chi.Router) {
			// r.With(middlewares.JWT_Auth).Post("/", h.LinkHandler.GenerateShortenedUrl)
			r.Post("/", h.LinkHandler.GenerateShortenedUrl)
		})

		//Route for redirecting
		r.Route("/{shortened_Url}", func(r chi.Router) {
			r.Get("/", h.LinkHandler.RedirectUrl)
		})
	})
	return r
}
