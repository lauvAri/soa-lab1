package model

// MaterialStatsByType 表示某个分类下的物资统计
type MaterialStatsByType struct {
	MaterialTypeID   int64  `gorm:"column:material_type_id" json:"material_type_id"`
	MaterialTypeName string `gorm:"column:material_type_name" json:"material_type_name"`
	TotalCount       int64  `gorm:"column:total_count" json:"total_count"`
	AvailableCount   int64  `gorm:"column:available_count" json:"available_count"`
}

// MaterialStats 汇总物资统计
type MaterialStats struct {
	ByType         []MaterialStatsByType `json:"by_type"`
	AvailableTotal int64                 `json:"available_total"`
}
