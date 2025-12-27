package settings

import (
	"time"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type ServerConfig struct {
	Addr               string
	Handler            *chi.Mux
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	IdleTimeout        time.Duration
	CORSAllowedOrigins []string
}

func loadServerConfig() ServerConfig {
	var tempArr []string
	router := CreateRouter()
	s := ServerConfig{
		Addr:               ":8080",
		Handler:            router,
		ReadTimeout:        time.Second * 2,
		WriteTimeout:       time.Second * 2,
		IdleTimeout:        time.Second * 2,
		CORSAllowedOrigins: tempArr,
	}
	return s
}


func CreateRouter() *chi.Mux {

	r := chi.NewRouter()

	//add middleware here
	r.Use(middleware.Logger)

	r.Route("/", func(r chi.Router) {

		//Routes for user mangement
		r.Route("/user", func(r chi.Router) {
			r.Post("/", CeateUser)
			r.Patch("/{userID}", UpdateUser)
			r.Delete("/{userId}", DeleteUser)
		})

		//Route for shortening
		r.Route("/shorten", func(r chi.Router) {
			r.Post("/", GenerateShortenedUrl)
		})

		//Route for redirecting
		r.Route("/{shortened_Url}", func(r chi.Router) {
			r.Get("/", RedirectUrl)
		})
	})
	return r
}
