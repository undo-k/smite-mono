package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

const grpcPort = "5002"

type Config struct {
	InProduction bool
}

func main() {
	log.SetReportCaller(true)
	app := Config{}

	setDebugMode(&app)

	app.gRPCListen()
}

func setDebugMode(app *Config) {
	if os.Getenv("DEBUG_MODE") == "False" {
		app.InProduction = true
	} else {
		app.InProduction = false
	}
}
