package main

import (
	"log/slog"
	"net/http"
	"os"

	"cryptomonitor/internal/web"
)

func main() {
	coins := []string{"bitcoin", "ethereum"}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	server, err := web.NewCryptoMonitorServer(http.DefaultClient, coins, logger)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	logger.Info("inicidando servidor", "addr", ":9000")

	err = http.ListenAndServe(":9000", server)
	logger.Error(err.Error())
	os.Exit(1)
}
