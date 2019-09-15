package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"models"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/inventory", getInventory).Methods("GET")
	router.HandleFunc("/inventory/{id:[0-9]+}", getInventory).Methods("GET")
	router.HandleFunc("/inventory/{id:[0-9]+}", updateInventory).Methods("PUT")
	router.HandleFunc("/inventory/{id:[0-9]+}", createInventory).Methods("POST")
	router.HandleFunc("/inventory/{id:[0-9]+}", deleteInventory).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}

func getInventory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inventoryId := params["id"]

	if inventoryId == "" {
		inventories := models.GetAllInventories()
		jsonResponse(w, r, inventories)
		return
	}

	id, _ := strconv.ParseUint(inventoryId, 10, 64)
	inventory := models.GetInventory(id)
	jsonResponse(w, r, inventory)
}

func updateInventory(w http.ResponseWriter, r *http.Request) {
}

func createInventory(w http.ResponseWriter, r *http.Request) {
}

func deleteInventory(w http.ResponseWriter, r *http.Request) {
}

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func jsonResponse(w http.ResponseWriter, r *http.Request, responseData interface{}) {
	setHeader(w)
	json.NewEncoder(w).Encode(responseData)
}
