package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type InventoryItem struct {
	ItemID      string `json:"item_id"`
	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
	WarehouseID string `json:"warehouse_id"`
}

var inventory []InventoryItem

func init() {
	inventory = loadInventory()
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	dbinit := NewDB("inventory.json")
	serviceinit := NewService(dbinit)
	controller := NewController(serviceinit)

	e.GET("/items/:id", controller.GetItem)
	e.POST("/items", controller.AddItem)
	e.PUT("/items/:id", controller.UpdateItem)
	e.DELETE("/items/:id", controller.DeleteItem)
	e.GET("/inventory/:warehouse_id", controller.GetInventoryByWarehouse)

	e.Logger.Fatal(e.Start(":8080"))
}

func loadInventory() []InventoryItem {

	return inventory
}
