package http

import (
	"context"
	"net"
	"net/http"

	"go.uber.org/fx"
)

func NewServer(lc fx.Lifecycle, mux *http.ServeMux) *http.Server {
	server := &http.Server{Addr: ":8080", Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			listener, err := net.Listen("tcp", ":8080") // TODO: Make these configurable.
			if err != nil {
				return err
			}
			go func() {
				err := server.Serve(listener)
				if err != nil {
					panic(err) // TODO: Get rid of this
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
	return server
}
