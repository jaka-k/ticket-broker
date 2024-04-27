package main

import "os"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	server := NewAPIServer(":" + port)
	server.Run()
}
