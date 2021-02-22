package structs

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Orders struct {
	gorm.Model
	Order_id      uint      `gorm:"primaryKey;autoIncrement:true`
	Customer_name string    `json:"costumer_name"`
	ordered_at    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"ordered_at"`
}

// func (Orders) TableName() string {
//     return "orders"
// }
