package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"models"

	"github.com/gorilla/mux"
)

type errorResponse struct {
	Message string
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/inventory", getInventory).Methods("GET")
	router.HandleFunc("/inventory", createInventory).Methods("POST")
	router.HandleFunc("/inventory/{id:[0-9]+}", getInventory).Methods("GET")
	router.HandleFunc("/inventory/{id:[0-9]+}", updateInventory).Methods("PUT")
	router.HandleFunc("/inventory/{id:[0-9]+}", deleteInventory).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}

func getInventory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inventoryId := params["id"]

	if inventoryId == "" {
		inventories := models.GetAllInventories()
		jsonResponse(w, r, inventories, 200)
		return
	}

	id, _ := strconv.ParseUint(inventoryId, 10, 64)
	inventory := models.GetInventory(id)

	if inventory == nil {
		jsonErrorResponse(w, r, 404, "Inventory not found")
		return
	}

	jsonResponse(w, r, inventory, 200)
}

func updateInventory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inventoryId := params["id"]
	var inventory models.Inventory
	_ = json.NewDecoder(r.Body).Decode(&inventory)
	id, _ := strconv.ParseUint(inventoryId, 10, 64)
	inventory.Id = id
	updatedInventory := models.UpdateInventory(inventory)
	jsonResponse(w, r, updatedInventory, 200)
}

func createInventory(w http.ResponseWriter, r *http.Request) {
	var inventory models.Inventory
	_ = json.NewDecoder(r.Body).Decode(&inventory)
	newInventoty := models.CreateInventory(inventory)
	jsonResponse(w, r, newInventoty, 201)
}

func deleteInventory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inventoryId := params["id"]
	id, _ := strconv.ParseUint(inventoryId, 10, 64)
	models.DeleteInventory(id)
	jsonResponse(w, r, nil, 200)
}

func setHeader(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
}

func jsonErrorResponse(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	setHeader(w, statusCode)
	json.NewEncoder(w).Encode(errorResponse{Message: message})
}

func jsonResponse(w http.ResponseWriter, r *http.Request, responseData interface{}, statusCode int) {
	setHeader(w, statusCode)
	json.NewEncoder(w).Encode(responseData)
}
