package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
)

type DB struct {
	inventoryFile string
}

func NewDB(inventoryFile string) *DB {
	return &DB{inventoryFile: inventoryFile}
}

func (db *DB) GetItem(itemID string) (*InventoryItem, error) {
	inventory, err := db.loadInventory()
	if err != nil {
		return nil, err
	}
	for _, item := range inventory {
		if item.ItemID == itemID {
			return &item, nil
		}
	}
	return nil, errors.New("Item not found")
}

func (db *DB) DeleteItem(itemID string) error {
	inventory, err := db.loadInventory()
	if err != nil {
		return err
	}
	for i, item := range inventory {
		if item.ItemID == itemID {
			inventory = append(inventory[:i], inventory[i+1:]...)
			return db.saveInventory(inventory)
		}
	}
	return errors.New("Item not found")
}

func (db *DB) AddItem(item InventoryItem) ([]InventoryItem, *InventoryItem, error) {
	inventory, err := db.loadInventory()
	if err != nil {
		return nil, nil, err
	}
	// Find the maximum item ID
	maxItemID := 0
	for _, existingItem := range inventory {
		itemIDInt, err := strconv.Atoi(existingItem.ItemID)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid item ID: %s", existingItem.ItemID)
		}
		if itemIDInt > maxItemID {
			maxItemID = itemIDInt
		}
	}
	// Increment the item ID and assign it to the new item
	item.ItemID = strconv.Itoa(maxItemID + 1)
	inventory = append(inventory, item)
	err = db.saveInventory(inventory)
	if err != nil {
		return nil, nil, err
	}
	return inventory, &item, nil
}

func (db *DB) saveInventory(inventory []InventoryItem) error {
	data, err := json.MarshalIndent(inventory, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(db.inventoryFile, data, 0644)
}

func (db *DB) UpdateItem(itemID string, updatedItem InventoryItem) error {
	inventory, err := db.loadInventory()
	if err != nil {
		return err
	}
	for i, item := range inventory {
		if item.ItemID == itemID {
			inventory[i] = updatedItem
			return db.saveInventory(inventory)
		}
	}
	return errors.New("Item not found")
}

func (db *DB) GetInventoryByWarehouse(warehouseID string) ([]InventoryItem, error) {
	inventory, err := db.loadInventory()
	if err != nil {
		return nil, err
	}
	var filteredInventory []InventoryItem
	for _, item := range inventory {
		if item.WarehouseID == warehouseID {
			filteredInventory = append(filteredInventory, item)
		}
	}
	return filteredInventory, nil
}

func (db *DB) loadInventory() ([]InventoryItem, error) {
	data, err := ioutil.ReadFile(db.inventoryFile)
	if err != nil {
		return nil, err
	}
	var inventory []InventoryItem
	if err := json.Unmarshal(data, &inventory); err != nil {
		return nil, err
	}
	return inventory, nil
}
