package field

import "gorm.io/gorm"

type BaseModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt Time
	UpdatedAt Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
