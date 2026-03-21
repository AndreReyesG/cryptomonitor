package main

import (
	"log"
	"net/http"

	"cryptomonitor/internal/web"
)

func main() {
	server, err := web.NewCryptoMonitorServer(http.DefaultClient, []string{"bitcoin", "ethereum"})
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Iniciando servidor en el puerto :9000")
	log.Fatal(http.ListenAndServe(":9000", server))
}
