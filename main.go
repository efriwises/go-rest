package main

import (
	"fmt"
	"go-rest/config"
	"go-rest/routes"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()

	r := mux.NewRouter()
	routes.RouteIndex(r)

	log.Println("Server running on port", config.ENV.PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), r)
}
