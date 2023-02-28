package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct {
	webPort int32
}

func main() {
	app := Config{
		webPort: 8090,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.webPort),
		Handler: app.routes(),
	}
	log.Println("Starting server at port %d", app.webPort)
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic("failed to start Server")
	}
}
