package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	config "github.com/spf13/viper"
)

type Server struct {
	router chi.Router
	server *http.Server
}

func NewServer() (*Server error) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Log Requests
	// if config.GetBool("server.log_requests") {
	// 	switch config.GetString("logger.encoding") {
	// 	case "stackdriver":
	// 		r.Use(loggerHTTPMiddlewareStackdriver(config.GetBool("server.log_requests_body"), config.GetStringSlice("server.log_disabled_http")))
	// 	default:
	// 		r.Use(loggerHTTPMiddlewareDefault(config.GetBool("server.log_requests_body"), config.GetStringSlice("server.log_disabled_http")))
	// 	}
	// }

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   config.GetStringSlice("server.cors.allowed_origins"),
		AllowedMethods:   config.GetStringSlice("server.cors.allowed_methods"),
		AllowedHeaders:   config.GetStringSlice("server.cors.allowed_headers"),
		AllowCredentials: config.GetBool("server.cors.allowed_credentials"),
		MaxAge:           config.GetInt("server.cors.max_age"),
	}).Handler)

	s := &Server {
		router: r
	}

	return s, nil
}

func (s *Server) Listen() error {
	
	http.Handle("/", s.router)

	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	} else {
		return nil 
	}
}

func (s *Server) Router() chi.Router {
	return s.router
}