package main

import (
	log "github.com/gofiber/fiber/v2/log"

	"github.com/luizarnoldch/mesapara2_backend/app"
)

func main() {
	log.Info("Server Starting")
	app.Start()
	log.Info("Server Finished")
}