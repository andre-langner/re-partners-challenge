package infrastructure

import (
	"errors"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// RunHTTPServer ...
func RunHTTPServer(log *zap.Logger, handler http.Handler) {
	srv := &http.Server{
		Addr:              ":3000",
		Handler:           handler,
		ReadHeaderTimeout: time.Second * 5,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("cannot run api server", zap.Error(err))
			panic(err)
		}
	}()
}
