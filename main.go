package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello world")

	godotenv.Load(".env")

	portSrting := os.Getenv("PORT")
	if portSrting == "" {
		log.Fatal("PORT is not found in the environment")
	}
	fmt.Println("Port:", portSrting)
}
