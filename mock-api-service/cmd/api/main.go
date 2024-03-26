package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8081"

type Config struct {
}

func main() {
	app := Config{}

	log.Printf("Starting mock api service on port %s\n", webPort)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
