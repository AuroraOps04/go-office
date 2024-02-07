package field

import "gorm.io/gorm"

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt Time           `                  json:"createdAt"`
	UpdatedAt Time           `                  json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"      json:"deletedAt"`
}
