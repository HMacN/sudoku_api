package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sudoku_api/config"
	"sudoku_api/config/config_key"

	"go.uber.org/fx"
)

func NewServer(lc fx.Lifecycle, mux *http.ServeMux) *http.Server {
	addr := fmt.Sprintf(":%d", config.Get(config_key.Port))
	server := &http.Server{Addr: addr, Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			listener, err := net.Listen("tcp", addr)
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
