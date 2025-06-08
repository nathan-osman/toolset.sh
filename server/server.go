package server

import (
	"context"
	"embed"
	"errors"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/nathan-osman/toolset.sh/manager"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	//go:embed static
	staticFS embed.FS
)

// Server provides the web interface.
type Server struct {
	server  http.Server
	logger  zerolog.Logger
	manager *manager.Manager
}

// New create a new Server instance.
func New(addr string, m *manager.Manager) (*Server, error) {
	var (
		r = gin.New()
		s = &Server{
			server: http.Server{
				Addr:    addr,
				Handler: r,
			},
			logger:  log.With().Str("package", "server").Logger(),
			manager: m,
		}
	)

	// Add the index page
	r.GET("/", s.index)

	// Page for running tools
	r.GET("/:name", s.tool)

	// Static files
	f, err := static.EmbedFolder(staticFS, "static")
	if err != nil {
		return nil, err
	}
	r.Use(static.Serve("/static", f))

	// Listen for connections in a separate goroutine
	go func() {
		defer s.logger.Info().Msg("server stopped")
		s.logger.Info().Msg("server started")
		if err := s.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error().Msg(err.Error())
		}
	}()

	return s, nil
}

// Close shuts down the server.
func (s *Server) Close() {
	s.server.Shutdown(context.Background())
}
