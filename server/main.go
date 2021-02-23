package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/greeting-svc/log"
	"github.com/greeting-svc/router"
)

func main() {

	os.Setenv("SERVICE_NAME", "greeting-svc")

	router := router.NewRouter()

	log.Logger().Info(fmt.Sprintf("Starting server on port: %v", 8080))

	log.Logger().Fatal(http.ListenAndServe(":8080", router))
}
