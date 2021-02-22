package controllers

import (
	"Assigment2/config"
	"Assigment2/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCurrentOrder(c *gin.Context) {

}

func CreateOrder(c *gin.Context) {

	var (
		order  structs.Orders
		item   structs.Items
		result gin.H
	)

	// id := c.PostForm("Id")
	customername := c.PostForm("costumer_name")
	description := c.PostForm("description")
	item_code := c.PostForm("item_code")
	quantity := c.PostForm("quantity")

	order.Customer_name = customername
	item.Description = description
	item.Item_code, _ = strconv.Atoi(item_code)
	item.Quantity, _ = strconv.Atoi(quantity)

	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record can't create!"})
		return
	}

	result = gin.H{
		"result": order,
		"item":   item,
	}
	c.JSON(http.StatusOK, result)
}

//update
func UpdateOrder(c *gin.Context) {
	id := c.Params.ByName("id")

	var (
		// order    structs.Orders
		// newOrder structs.Orders
		item    structs.Items
		newItem structs.Items

		result gin.H
	)

	if err := config.DB.Where("id = ?", id).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record Not Found",
		})
		return
	}

	// customername := c.PostForm("costumer_name")
	description := c.PostForm("description")
	item_code := c.PostForm("item_code")
	quantity := c.PostForm("quantity")

	// newOrder.Customer_name = customername
	newItem.Description = description
	newItem.Item_code, _ = strconv.Atoi(item_code)
	newItem.Quantity, _ = strconv.Atoi(quantity)

	if err := config.DB.Model(&item).Updates(newItem).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record can't update!",
		})

		c.JSON(http.StatusOK, result)
	}
}
