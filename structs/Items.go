package structs

type Items struct {
	// gorm.Model
	//GORM defined a gorm.Model struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	Item_id     int    `gorm:"primary_key"`
	Item_code   int    `gorm:"type:int"`
	Description string `gorm:"type:varchar(200)`
	Quantity    int    `gorm:"type:int`
	Order_id    int    `gorm:"type:int`

	// Order    []Orders `gorm:"foreignKey:o_id;association_foreignkey:Order_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	// Order    []Orders `gorm:"foreignKey:o_id;association_foreignkey:Order_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

// func (Items) TableName() string {
// 	return "Items"
// }
