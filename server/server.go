package server

import (
	"context"
	// "fmt"
	"github.com/dgraph-io/dgo/v200"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/roadev/goapi/controllers"
	"github.com/roadev/goapi/database"
	config "github.com/spf13/viper"
	"net/http"
)

type Server struct {
	router       chi.Router
	server       *http.Server
	dgraphClient *dgo.Dgraph
	ctx          context.Context
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

	dgraphClient, ctx := database.NewDatabaseConnection()

	s := &Server{
		router:       r,
		dgraphClient: dgraphClient,
		ctx:          ctx,
	}

	return s, nil
}

func (s *Server) Router() chi.Router {
	s.router.Get("/buyers", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetAllBuyers(s.dgraphClient, s.ctx, w)
	})

	s.router.Get("/load_buyers", func(w http.ResponseWriter, r *http.Request) {
		controllers.LoadBuyers(1600053936468)
	})

	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))

	})

	return s.router
}

func Listen(r chi.Router) {

	http.ListenAndServe(":3000", r)

}
