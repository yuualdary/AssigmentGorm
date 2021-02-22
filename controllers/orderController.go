package controllers

import (
	"Assigment2/config"
	"Assigment2/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCurrentOrderById(c *gin.Context) {
	var (
		id     = c.Params.ByName("id")
		order  []structs.Orders
		item   []structs.Items
		result gin.H
	)

	// if err := config.DB.Where("order_id = ?", id).First(&order).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Record not found!",
	// 	})

	// } else {
	// 	result = gin.H{
	// 		"Item":  item,
	// 		"Order": order,
	// 		"count": 1,
	// 	}

	// }

	if err := config.DB.Joins("left join items on items.order_id = orders.id").Where("order_id = ?", id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})

	} else {
		result = gin.H{
			"Item":  item,
			"Order": order,
			"count": 1,
		}

	}
	c.JSON(http.StatusOK, result)

}

func CreateOrder(c *gin.Context) {

	var (
		order  structs.Orders
		item   structs.Items
		result gin.H
	)

	// id := c.PostForm("Id")
	customer_name := c.PostForm("customer_name")
	description := c.PostForm("description")
	item_code := c.PostForm("item_code")
	getCode, _ := strconv.Atoi(item_code)
	quantity := c.PostForm("quantity")
	getQty, _ := strconv.Atoi(quantity)

	order.Customer_name = customer_name
	item.Description = description
	item.Item_code = getCode
	item.Quantity = getQty

	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error Order": "Record can't create!",
		})
	}

	if getr := config.DB.Create(&item).Error; getr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error Item": "Record can't create!",
		})
	} else {
		result = gin.H{
			"order": order,
			"item":  item,
		}
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
