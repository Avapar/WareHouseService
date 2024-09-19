package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) GetItem(ctx echo.Context) error {
	itemID := ctx.Param("id")
	item, err := c.service.GetItem(itemID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "Item not found")
	}
	return ctx.JSON(http.StatusOK, item)
}

func (c *Controller) AddItem(ctx echo.Context) error {
	var newItem InventoryItem
	if err := ctx.Bind(&newItem); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request body")
	}
	newer, err := c.service.AddItem(newItem)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to add item")
	}
	return ctx.JSON(http.StatusCreated, newer)
}

func (c *Controller) UpdateItem(ctx echo.Context) error {
	itemID := ctx.Param("id")
	var updatedItem InventoryItem
	if err := ctx.Bind(&updatedItem); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request body")
	}
	err := c.service.UpdateItem(itemID, updatedItem)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to update item")
	}
	return ctx.JSON(http.StatusOK, "Item updated successfully")
}

func (c *Controller) DeleteItem(ctx echo.Context) error {
	itemID := ctx.Param("id")
	err := c.service.DeleteItem(itemID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to delete item")
	}
	return ctx.JSON(http.StatusOK, "Item deleted successfully")
}

func (c *Controller) GetInventoryByWarehouse(ctx echo.Context) error {
	warehouseID := ctx.Param("warehouse_id")
	inventory, err := c.service.GetInventoryByWarehouse(warehouseID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to retrieve inventory")
	}
	return ctx.JSON(http.StatusOK, inventory)
}
