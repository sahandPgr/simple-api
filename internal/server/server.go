package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ServerConfig struct {
	server_url  string
	server_port string
}

type IServer interface {
	InitialServer() error
}

// Initializing the server config
func NewServer(url string, port string) ServerConfig {
	var serverConfig = ServerConfig{
		server_url:  url,
		server_port: port,
	}
	return serverConfig
}

func (sc ServerConfig) InitialServer(r *mux.Router) error {
	//Listening and serve
	serverUrl := sc.server_url + ":" + sc.server_port
	fmt.Printf("Listening on %s...\n", serverUrl)
	log.Fatal(http.ListenAndServe(serverUrl, r))
	return nil
}
