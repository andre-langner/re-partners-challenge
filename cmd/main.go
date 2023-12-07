package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"re_partners/internal/handlers/calculatepackshandler"
	"re_partners/internal/handlers/healthhandler"
	"re_partners/internal/infrastructure"
	"re_partners/internal/repository/packsrepo"
	"re_partners/internal/services/calculatepacksservice"
)

func main() {
	logger, err := infrastructure.LoadLogger()
	if err != nil {
		panic(fmt.Sprintf("could not create logger: %s", err.Error()))
	}
	defer func() { _ = logger.Sync() }()

	repo := packsrepo.New()

	service := calculatepacksservice.New(repo)

	router := infrastructure.LoadRouter(
		healthhandler.New(),
		calculatepackshandler.New(service),
	)

	infrastructure.RunHTTPServer(logger, router)

	logger.Info("Running....")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sigs
}
