package model

import (
	"time"
)

// Material 映射到表 `materials_info`
type Material struct {
	MaterialID          int64      `gorm:"column:material_id;primaryKey;autoIncrement" json:"materialId"`
	MaterialName        string     `gorm:"column:material_name;type:varchar(255);not null" json:"materialName"`
	MaterialTypeID      int64      `gorm:"column:material_type_id;not null" json:"materialTypeId"`
	MaterialDesc        *string    `gorm:"column:material_desc;type:varchar(255)" json:"materialDesc,omitempty"`
	MaterialStatus      int16      `gorm:"column:material_status;type:smallint;not null;default:0" json:"materialStatus"`
	MaterialPurchasedAt *time.Time `gorm:"column:material_purchased_at;type:datetime" json:"materialPurchasedAt,omitempty"`
	MaterialLocation    *string    `gorm:"column:material_location;type:varchar(100)" json:"materialLocation,omitempty"`
}

// TableName 指定 GORM 使用的表名
func (Material) TableName() string { return "materials_info" }
