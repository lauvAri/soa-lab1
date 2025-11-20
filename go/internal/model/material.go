package model

import (
	"encoding/json"
	"time"
)

// Material 映射到表 `materials_info`
type Material struct {
	MaterialID          int64      `gorm:"column:material_id;primaryKey;autoIncrement" json:"material_id"`
	MaterialName        string     `gorm:"column:material_name;type:varchar(255);not null" json:"material_name"`
	MaterialTypeID      int64      `gorm:"column:material_type_id;not null" json:"material_type_id"`
	MaterialDesc        *string    `gorm:"column:material_desc;type:varchar(255)" json:"material_desc,omitempty"`
	MaterialStatus      int16      `gorm:"column:material_status;type:smallint;not null;default:0" json:"material_status"`
	MaterialPurchasedAt *time.Time `gorm:"column:material_purchased_at;type:datetime" json:"material_purchased_at,omitempty"`
	MaterialLocation    *string    `gorm:"column:material_location;type:varchar(100)" json:"material_location,omitempty"`
}

// TableName 指定 GORM 使用的表名
func (Material) TableName() string { return "materials_info" }

// UnmarshalJSON 兼容旧的 camelCase 字段，确保可以读取 materialId/materialName 等字段
func (m *Material) UnmarshalJSON(data []byte) error {
	type alias Material
	var converted alias
	if err := json.Unmarshal(data, &converted); err != nil {
		return err
	}
	*m = Material(converted)

	var legacy struct {
		MaterialID          *int64     `json:"materialId"`
		MaterialName        *string    `json:"materialName"`
		MaterialTypeID      *int64     `json:"materialTypeId"`
		MaterialDesc        *string    `json:"materialDesc"`
		MaterialStatus      *int16     `json:"materialStatus"`
		MaterialPurchasedAt *time.Time `json:"materialPurchasedAt"`
		MaterialLocation    *string    `json:"materialLocation"`
	}
	if err := json.Unmarshal(data, &legacy); err != nil {
		return err
	}
	if legacy.MaterialID != nil {
		m.MaterialID = *legacy.MaterialID
	}
	if legacy.MaterialName != nil {
		m.MaterialName = *legacy.MaterialName
	}
	if legacy.MaterialTypeID != nil {
		m.MaterialTypeID = *legacy.MaterialTypeID
	}
	if legacy.MaterialDesc != nil {
		m.MaterialDesc = legacy.MaterialDesc
	}
	if legacy.MaterialStatus != nil {
		m.MaterialStatus = *legacy.MaterialStatus
	}
	if legacy.MaterialPurchasedAt != nil {
		m.MaterialPurchasedAt = legacy.MaterialPurchasedAt
	}
	if legacy.MaterialLocation != nil {
		m.MaterialLocation = legacy.MaterialLocation
	}
	return nil
}
