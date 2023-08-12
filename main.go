package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

    "github.com/charmbracelet/log"
	"github.com/elias-gill/poliapi/src/routers"
	"github.com/elias-gill/poliapi/src/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/joho/godotenv/autoload"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Api para posible politerminal web
// @version 1.0
// @description API sencilla para manejar los horarios de la facultad
// @contact.name PoliAPI
// @contact.email eliasgill42@gmail.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host
// @BasePath /
func main() {
	// instantiate chi router
	r := chi.NewRouter()

	// middlewares
	r.Use(requestSizeLimiter(1 << 19)) // maximo de 1MB
	r.Use(middleware.Logger)
	r.Use(utils.JwtMidleware)
	r.Mount("/swagger", httpSwagger.WrapHandler)

	// routes
	r.Route("/user", routers.UsersHandler)

	// configure server to run
	srv := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: r,
	}
	// Using a channel to detect when a shutdown signal has been received.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	// function to gracefully shutdown
	go func() {
		<-sigCh
		log.Print("Shutdown signal received, shutting down gracefully...")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Error during shutdown: %v\n", err)
		}
		log.Print("Server has been shutdown.")
	}()

	// start the server
	log.Print("Server listening at port", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", srv.Addr, err)
	}
}

// middleware to limit the body size of a request
func requestSizeLimiter(limit int64) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Body = http.MaxBytesReader(w, r.Body, limit)

			buf := new(strings.Builder)
			_, err := io.Copy(buf, r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if int64(buf.Len()) > limit {
				http.Error(w, "Request size exceeds limit", http.StatusRequestEntityTooLarge)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
