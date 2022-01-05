package main

import (
	"net/http"
	"github.com/Takachii/amanah-int/config"
	"github.com/Takachii/amanah-int/model"
  "strconv"
	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDB()
	route := echo.New()
	route.POST("inventory/add", func(c echo.Context) error {
		item := new(model.Inventories)
		c.Bind(item)
		response := new(Response)
		if item.CreateItem() != nil { // method create user
			response.ErrorCode = 10
			response.Message = "Gagal create data user"
		} else {
			response.ErrorCode = 0
			response.Message = "Sukses create data user"
			response.Data = *item
		}
		return c.JSON(http.StatusOK, response)
	})

	route.PUT("inventory/update/:id", func(c echo.Context) error {
		item := new(model.Inventories)
		c.Bind(item)
		response := new(Response)
    sID , _ := strconv.Atoi(c.Param("id"))
		if item.UpdateItem(sID) != nil { // method update user
			response.ErrorCode = 10
			response.Message = "Gagal update data user"
		} else {
			response.ErrorCode = 0
			response.Message = "Sukses update data user"
			response.Data = *item
		}
		return c.JSON(http.StatusOK, response)
	})

	route.DELETE("inventory/delete/:id", func(c echo.Context) error {
    sID , _ := strconv.Atoi(c.Param("id"))
		item, _ := model.GetOneById(sID) // method get by email
		response := new(Response)

		if item.DeleteItem() != nil { // method update user
			response.ErrorCode = 10
			response.Message = "Gagal menghapus data user"
		} else {
			response.ErrorCode = 0
			response.Message = "Sukses menghapus data user"
		}
		return c.JSON(http.StatusOK, response)
	})

	route.GET("inventory/search", func(c echo.Context) error {
		response := new(Response)
		items, err := model.GetAll() // method get all
		if err != nil {
			response.ErrorCode = 10
			response.Message = "Gagal melihat data user"
		} else {
			response.ErrorCode = 0
			response.Message = "Sukses melihat data user"
			response.Data = items
		}
		return c.JSON(http.StatusOK, response)
	})

  route.PUT("purchase/:id/:val", func(c echo.Context) error {
    return nil
  })

	route.Start(":9000")
}

type Response struct {
	ErrorCode int         `json:"error_code" form:"error_code"`
	Message   string      `json:"message" form:"message"`
	Data      interface{} `json:"data"`
}
