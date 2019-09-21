package models

import (
	"database/sql"
)

// Inventory inventory type
type Inventory struct {
	Id                          uint64
	VehicleIdentificationNumber string
	ModelName                   string
	Producer                    string
	Year                        uint16
	MSRP                        float64
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
func GetAllInventories() (interface{}, error) {
	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM inventory")
	defer rows.Close()

	inventories := make([]Inventory, 0)
	for rows.Next() {
		inventory, _ := getInventoryFromRows(rows)
		inventories = append(inventories, inventory)
	}

	if err != nil {
		return nil, err
	}

	return inventories, nil
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
func DeleteInventory(id uint64) error {
	db := connect()
	defer db.Close()

	result, err := db.Query("DELETE FROM inventory WHERE id=?", id)
	result.Close()

	return err
}

// UpdateInventory update an inventory exceptions or restrictions can be added but for now you can edit however you want
func UpdateInventory(inventoryUpdates Inventory) (interface{}, error) {
	db := connect()
	defer db.Close()

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
	defer result.Close()

	// should do something meaningful then just panic
	if error != nil {
		return nil, error
	}

	return GetInventory(inventoryUpdates.Id), nil
}

// CreateInventory create an inventory
func CreateInventory(inventory Inventory) (interface{}, error) {
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

	if error != nil {
		return nil, error
	}

	row := db.QueryRow("SELECT * FROM inventory ORDER BY Id DESC LIMIT 1")
	newInventory, _ := getInventoryFromRow(row)
	return newInventory, nil
}
