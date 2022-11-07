package main

import (
	
	"log"
	"os"
	"restapiexample/src/configs/command"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error while loading .env file")
	}

	if err := command.Run(os.Args[1:]); err != nil {
		log.Fatal(err.Error())
	}
}