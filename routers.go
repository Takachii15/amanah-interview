package main

import (
	"github.com/Takachii/amanah-int/handlers"
	"github.com/labstack/echo/v4"
)

func InventoryRouter() {
	routes := echo.New()
	route := routes.Group("/inventory")
	{
		route.POST("/inventory/add", handlers.AddInventory)
		route.PUT("/inventory/update/:id", handlers.UpdateInventory)
		route.DELETE("/inventory/delete/:id", handlers.DeleteInventory)
		route.GET("/inventory/search", handlers.GetInventories)
	}
	routes.Start(":9000")
}

type Response struct {
	ErrorCode int         `json:"error_code" form:"error_code"`
	Message   string      `json:"message" form:"message"`
	Data      interface{} `json:"data"`
}
