package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	fmt.Println("Pulse Mailer")
}
