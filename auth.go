package auth

import (
	"context"
	"net/http"
	"os"

	"github.com/smxlong/kit/logger"
	"github.com/smxlong/kit/service"
	"github.com/smxlong/kit/webserver"
	"github.com/spf13/cobra"
)

// Auth is the auth application.
type Auth struct {
	listenAddress string
}

// New creates a new auth application.
func New() *Auth {
	return &Auth{}
}

// Command returns the command for the auth application.
func (a *Auth) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "auth",
		Short:         "authentication server",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			a.bindEnv()
			return service.Run(func(ctx context.Context, l logger.Logger) error {
				return a.Run(ctx, l)
			})
		},
	}
	return cmd
}

// bindEnv binds environment variables to the auth application.
func (a *Auth) bindEnv() {
	a.listenAddress = os.Getenv("LISTEN_ADDRESS")
	if a.listenAddress == "" {
		a.listenAddress = ":8080"
	}
}

// Run runs the auth application.
func (a *Auth) Run(ctx context.Context, l logger.Logger) error {
	return webserver.ListenAndServe(ctx, &http.Server{
		Addr:    a.listenAddress,
		Handler: a.routes(),
	})
}
