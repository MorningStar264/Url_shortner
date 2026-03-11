package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/MorningStar264/Url_shortner/internal/config"
	"github.com/MorningStar264/Url_shortner/internal/handler"
	"github.com/MorningStar264/Url_shortner/internal/repository"
	"github.com/MorningStar264/Url_shortner/internal/router"
	"github.com/MorningStar264/Url_shortner/internal/server"
)

const DefaultContextTimeout = 30

func main() {
	// Loading the configs form .env
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	// Initialize server
	srv, err := server.New(cfg)
	if err != nil {
		log.Fatal("failed to initialize server", err)
	}

	// Initialize repositories, services, and handlers
	
	repos := repository.NewRepositories(srv)

	handlers := handler.NewHandlers(srv, repos)

	// Initialize router
	r := router.NewRouter(srv, handlers)

	// Setup HTTP server
	srv.SetupHTTPServer(r)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)

	// Start server
	go func() {
		if err = srv.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("failed to start server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), DefaultContextTimeout*time.Second)

	if err = srv.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown")
	}
	stop()
	cancel()

	log.Println("server exited properly")
}
