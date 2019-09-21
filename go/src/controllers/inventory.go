package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"models"

	"github.com/gorilla/mux"
)

// SetInventoryRoute set inventory routers
func SetInventoryRoute(router *mux.Router) {
	router.HandleFunc("/inventory", getInventory).Methods(http.MethodGet)
	router.HandleFunc("/inventory", createInventory).Methods(http.MethodPost)
	router.HandleFunc("/inventory/{id:[0-9]+}", getInventory).Methods(http.MethodGet)
	router.HandleFunc("/inventory/{id:[0-9]+}", updateInventory).Methods(http.MethodPut)
	router.HandleFunc("/inventory/{id:[0-9]+}", deleteInventory).Methods(http.MethodDelete)
}

func getInventory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inventoryId := params["id"]

	if inventoryId == "" {
		inventories, e := models.GetAllInventories()

		if e != nil {
			JSONErrorResponse(w, r, 500, e.Error())
			return
		}

		JSONResponse(w, r, inventories, 200)
		return
	}

	id, e := strconv.ParseUint(inventoryId, 10, 64)

	if e != nil {
		JSONErrorResponse(w, r, 500, e.Error())
		return
	}

	inventory := models.GetInventory(id)

	if inventory == nil {
		JSONErrorResponse(w, r, 404, "Inventory not found")
		return
	}

	JSONResponse(w, r, inventory, 200)
}

func updateInventory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inventoryId := params["id"]
	var inventory models.Inventory
	_ = json.NewDecoder(r.Body).Decode(&inventory)
	id, e := strconv.ParseUint(inventoryId, 10, 64)

	if e != nil {
		JSONErrorResponse(w, r, 500, e.Error())
		return
	}

	inventory.Id = id
	updatedInventory, err := models.UpdateInventory(inventory)

	if err != nil {
		JSONErrorResponse(w, r, 500, err.Error())
		return
	}

	JSONResponse(w, r, updatedInventory, 200)
}

func createInventory(w http.ResponseWriter, r *http.Request) {
	var inventory models.Inventory
	_ = json.NewDecoder(r.Body).Decode(&inventory)
	newInventoty, err := models.CreateInventory(inventory)

	if err != nil {
		JSONErrorResponse(w, r, 500, err.Error())
		return
	}

	JSONResponse(w, r, newInventoty, 201)
}

func deleteInventory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inventoryId := params["id"]
	id, err := strconv.ParseUint(inventoryId, 10, 64)

	if err != nil {
		JSONErrorResponse(w, r, 500, err.Error())
		return
	}

	e := models.DeleteInventory(id)

	if e != nil {
		JSONErrorResponse(w, r, 500, e.Error())
		return
	}

	JSONResponse(w, r, nil, 200)
}
