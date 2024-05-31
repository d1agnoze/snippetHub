package main

import (
	"fmt"
	"snippetHub/internal/auth"
	"snippetHub/internal/server"
)

func main() {

  auth.NewAuth()
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
