package main

import (
	"os"

	"github.com/joaomarcuslf/keyforge-name-of-the-day/server"
)

func main() {
	server := server.NewServer(os.Getenv("PORT"))

	server.Run()
}
