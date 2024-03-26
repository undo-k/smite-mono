package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/undo-k/smite-mono/protos/protos"
)

const grpcPort = "5001"

type Config struct {
	GodCache     map[string]*protos.God
	InProduction bool
}

func main() {
	log.SetReportCaller(true)
	app := Config{}

	setDebugMode(&app)

	godCache, err := createGodCache()
	if err != nil {
		log.Error("Failed to create god cache:")
		log.Error(err)
	}
	app.GodCache = godCache

	app.gRPCListen()
}

func setDebugMode(app *Config) {
	if os.Getenv("DEBUG_MODE") == "False" {
		app.InProduction = true
	} else {
		app.InProduction = false
	}
}
