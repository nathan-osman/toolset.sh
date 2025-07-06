package server

import (
	"context"
	"embed"
	"errors"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	//go:embed static
	staticFS embed.FS
)

// Server provides the web interface.
type Server struct {
	server http.Server
	logger zerolog.Logger
}

// New create a new Server instance.
func New(addr string) (*Server, error) {
	var (
		r = gin.New()
		s = &Server{
			server: http.Server{
				Addr:    addr,
				Handler: r,
			},
			logger: log.With().Str("package", "server").Logger(),
		}
	)

	// Handle errors by rendering the error page
	r.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		msg := "an unknown error has occurred"
		switch v := err.(type) {
		case error:
			msg = v.Error()
		}
		s.sendError(c, msg)
		c.Abort()
	}))

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
