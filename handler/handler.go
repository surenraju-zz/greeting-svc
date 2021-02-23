package handler

import (
	"fmt"
	"net/http"

	"github.com/greeting-svc/log"

	"github.com/gorilla/mux"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	log.Logger().Info(fmt.Sprintf("Processing name: %v", name))

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, fmt.Sprintf("Hello %s \n", name))

}
