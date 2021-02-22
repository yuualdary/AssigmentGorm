package structs

import (
	"github.com/jinzhu/gorm"
)

type Items struct {
	gorm.Model
	Item_id     uint   `gorm:"primaryKey;autoIncrement:true;`
	Item_code   int    `json:item_code`
	Description string `json:description`
	Quantity    int    `json:quantity`
	// Orders      Orders `gorm:"references:Order_id`
	Order_id uint     `gorm:"column:o_id" json:"-"`
	Order    []Orders `gorm:"foreignKey:o_id;association_foreignkey:Order_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

// func (Items) TableName() string {
// 	return "Items"
// }
