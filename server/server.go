package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	config "github.com/spf13/viper"
)

type Server struct {
	router chi.Router
	server *http.Server
}

type Router struct{}

func NewServer() (*Server, error) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   config.GetStringSlice("server.cors.allowed_origins"),
		AllowedMethods:   config.GetStringSlice("server.cors.allowed_methods"),
		AllowedHeaders:   config.GetStringSlice("server.cors.allowed_headers"),
		AllowCredentials: config.GetBool("server.cors.allowed_credentials"),
		MaxAge:           config.GetInt("server.cors.max_age"),
	}).Handler)

	s := &Server{
		router: r,
	}

	return s, nil
}

func (s *Server) Router() chi.Router {
	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi1!!"))
	})

	s.router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	return s.router
}

func Listen(r chi.Router) {

	http.ListenAndServe(":3000", r)

}
