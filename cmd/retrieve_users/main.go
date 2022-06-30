package main

import (
	api "github.com/blacksfk/retrieve_users/http"
	"github.com/blacksfk/uf"
	"log"
	"net/http"
)

const (
	ADDRESS = "localhost:6060"
)

func main() {
	sconf := &uf.Config{
		ErrorLogger:  logStdout,
		AccessLogger: uf.LogStdout,
	}

	// create server
	server := uf.NewServer(sconf)

	// define routes
	routes(server)

	log.Printf("Starting server on %s\n", ADDRESS)

	// anchors aweigh!
	log.Fatal(http.ListenAndServe(ADDRESS, server))
}

// Error logger.
func logStdout(e error) {
	log.Println(e)
}

// Route definitions
func routes(server *uf.Server) {
	server.Post("/retrieveUsers", api.RetrieveUsers)
}
