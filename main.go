package main

import (
	"log"

	"github.com/luizarnoldch/mesapara2_backend/app"
)

func main() {
	log.Println("Server Starting")
	app.Start()
	log.Println("Server Finished")
}