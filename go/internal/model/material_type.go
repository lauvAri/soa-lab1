package model

// 种类表模型
type MaterialType struct {
	ID   uint   `gorm:"primaryKey;column:material_type_id"`
	Name string `gorm:"column:material_type_name;type:varchar(100);not null"`
}

func (MaterialType) TableName() string {
	return "material_type"
}