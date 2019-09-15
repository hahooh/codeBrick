package models

import (
	"database/sql"
	"fmt"
)

// Inventory inventory type
type Inventory struct {
	Id                          uint64
	VehicleIdentificationNumber string
	ModelName                   string
	Producer                    string
	Year                        uint16
	MSRP                        float32
	Status                      string
	Booked                      bool
	Listed                      bool
	CreatedDate                 string
	ModifiedDate                string
}

func getInventoryFromRows(rows *sql.Rows) (Inventory, error) {
	var inventory Inventory
	err := rows.Scan(&inventory.Id,
		&inventory.VehicleIdentificationNumber,
		&inventory.ModelName, &inventory.Producer,
		&inventory.Year, &inventory.MSRP,
		&inventory.Status, &inventory.Booked,
		&inventory.Listed,
		&inventory.CreatedDate,
		&inventory.ModifiedDate)

	if err != nil {
		return inventory, err
	}

	return inventory, nil
}

func getInventoryFromRow(row *sql.Row) (Inventory, error) {
	var inventory Inventory
	err := row.Scan(&inventory.Id,
		&inventory.VehicleIdentificationNumber,
		&inventory.ModelName, &inventory.Producer,
		&inventory.Year, &inventory.MSRP,
		&inventory.Status, &inventory.Booked,
		&inventory.Listed,
		&inventory.CreatedDate,
		&inventory.ModifiedDate)

	if err != nil {
		return inventory, err
	}

	return inventory, nil
}

// GetAllInventories return all inventories
func GetAllInventories() interface{} {
	db := connect()
	defer db.Close()

	rows, error := db.Query("SELECT * FROM inventory")
	defer rows.Close()

	inventories := make([]Inventory, 0)
	for rows.Next() {
		inventory, _ := getInventoryFromRows(rows)
		inventories = append(inventories, inventory)
	}

	// should do something meaningful then just panic
	if error != nil {
		fmt.Println("ERROR!!!")
		panic(error.Error())
	}

	return inventories
}

// GetInventory get an inventory by Id
func GetInventory(id uint64) interface{} {
	db := connect()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM inventory WHERE Id=?", id)
	inventory, err := getInventoryFromRow(row)
	if err != nil {
		return nil
	}
	return inventory
}

// DeleteInventory delete an inventory by Id
func DeleteInventory(id uint64) interface{} {
	db := connect()
	defer db.Close()

	result, error := db.Query("DELETE FROM inventory WHERE id=?", id)
	result.Close()

	// should do something meaningful then just panic
	if error != nil {
		panic(error.Error())
	}

	return result
}

// UpdateInventory update an inventory exceptions or restrictions can be added but for now you can edit however you want
func UpdateInventory(inventoryUpdates Inventory) interface{} {
	db := connect()
	defer db.Close()

	fmt.Println("Before UPDATE query")

	result, error := db.Query("UPDATE inventory SET VehicleIdentificationNumber=?, ModelName=?, Producer=?, Year=?, MSRP=?, Status=?, Booked=?, Listed=? WHERE Id=?",
		inventoryUpdates.VehicleIdentificationNumber,
		inventoryUpdates.ModelName,
		inventoryUpdates.Producer,
		inventoryUpdates.Year,
		inventoryUpdates.MSRP,
		inventoryUpdates.Status,
		inventoryUpdates.Booked,
		inventoryUpdates.Listed,
		inventoryUpdates.Id)

	fmt.Println("After UPDATE query")

	defer result.Close()

	fmt.Println("after defer", result, error)

	// should do something meaningful then just panic
	if error != nil {
		panic(error.Error())
	}

	return GetInventory(inventoryUpdates.Id)
}

// CreateInventory create an inventory
func CreateInventory(inventory Inventory) interface{} {
	db := connect()
	defer db.Close()

	_, error := db.Query("INSERT INTO inventory (VehicleIdentificationNumber, ModelName, Producer, Year, MSRP, Status, Booked, Listed) VALUE (?,?,?,?,?,?,?,?)",
		inventory.VehicleIdentificationNumber,
		inventory.ModelName,
		inventory.Producer,
		inventory.Year,
		inventory.MSRP,
		inventory.Status,
		inventory.Booked,
		inventory.Listed)

	// should do something meaningful then just panic
	if error != nil {
		panic(error.Error())
	}

	row := db.QueryRow("SELECT * FROM inventory ORDER BY Id DESC LIMIT 1")
	newInventory, _ := getInventoryFromRow(row)
	return newInventory
}
