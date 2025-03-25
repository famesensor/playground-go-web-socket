package infra

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/famesensor/go-template/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HTTPServer struct {
	cfg *config.Config
	e   *echo.Echo
}

type Handler interface {
	RegisterRoutes(e *echo.Echo)
}

func NewHTTPServer(cfg *config.Config) *HTTPServer {
	e := echo.New()
	e.Use(
		middleware.Recover(),
		middleware.Logger(),
		middleware.CORS(),
	)

	return &HTTPServer{
		cfg: cfg,
		e:   e,
	}
}

func (hs *HTTPServer) RegisterRoute(handlers ...Handler) {
	for _, handler := range handlers {
		handler.RegisterRoutes(hs.e)
	}
}

func (hs *HTTPServer) Run(cleanFuncs ...func()) {
	// Start server
	go func() {
		if err := hs.e.Start(":" + hs.cfg.App.Port); err != nil && err != http.ErrServerClosed {
			log.Fatal("shutting down the server")
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	for _, cleanup := range cleanFuncs {
		cleanup()
	}

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := hs.e.Shutdown(ctx); err != nil {
		hs.e.Logger.Fatal(err)
	}
}
