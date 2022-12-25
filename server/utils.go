package server

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

func graceServe(ctx context.Context, addr string, h http.Handler) error {
	addr = strings.TrimSpace(addr)
	if addr == "" {
		addr = ":8080"
	}

	srv := &http.Server{
		Addr:    addr,
		Handler: h,
	}

	errCh := make(chan error)
	go func() {
		defer close(errCh)

		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		graceCtx, cancel := context.WithTimeout(context.Background(), gracePeriod)
		defer cancel()
		return srv.Shutdown(graceCtx)

	case err := <-errCh:
		return err
	}
}
