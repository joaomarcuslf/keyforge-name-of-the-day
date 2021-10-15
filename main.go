package main

import (
	"os"

	"github.com/joaomarcuslf/keyforge-name-of-the-day/routines"
	"github.com/joaomarcuslf/keyforge-name-of-the-day/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	server := server.NewServer(os.Getenv("PORT"))
	cj := routines.NewCronJob()

	go cj.Run()

	server.Run()
}
