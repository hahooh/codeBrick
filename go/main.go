package main

import (
	"net/http"

	"controllers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{http.MethodOptions, http.MethodPost, http.MethodPut, http.MethodGet, http.MethodDelete})
	router.HandleFunc("*", handleOption).Methods(http.MethodOptions)

	controllers.SetInventoryRoute(router)

	http.ListenAndServe(":8000", handlers.CORS(allowedMethods, allowedOrigins)(router))
}

func handleOption(w http.ResponseWriter, r *http.Request) {
	controllers.JSONResponse(w, r, nil, 200)
}
