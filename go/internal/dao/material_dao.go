package dao

import (
	"context"

	"gorm.io/gorm"

	"materials-service/internal/model"
)

// CreateMaterial 在数据库中创建一条物资记录
func CreateMaterial(ctx context.Context, m *model.Material) error {
	return model.DB.WithContext(ctx).Create(m).Error
}

// GetMaterialByID 根据主键查询物资
func GetMaterialByID(ctx context.Context, id int64) (*model.Material, error) {
	var mat model.Material
	if err := model.DB.WithContext(ctx).First(&mat, "material_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &mat, nil
}

// UpdateMaterial 按主键更新整条记录（零值也会更新）
func UpdateMaterial(ctx context.Context, m *model.Material) error {
	return model.DB.WithContext(ctx).
		Model(&model.Material{}).
		Where("material_id = ?", m.MaterialID).
		Updates(m).Error
}

// DeleteMaterial 根据主键删除
func DeleteMaterial(ctx context.Context, id int64) error {
	result := model.DB.WithContext(ctx).Delete(&model.Material{}, "material_id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// ListMaterials 分页查询物资列表，返回列表与总数
func ListMaterials(ctx context.Context, offset, limit int) ([]model.Material, int64, error) {
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	var (
		items []model.Material
		total int64
	)

	db := model.DB.WithContext(ctx).Model(&model.Material{})
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Offset(offset).Limit(limit).Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// GetMaterialStats 统计每个分类的物资数量及可用数量
func GetMaterialStats(ctx context.Context) ([]model.MaterialStatsByType, error) {
	var stats []model.MaterialStatsByType
	err := model.DB.WithContext(ctx).
		Table("material_type AS mt").
		Select(`
			mt.material_type_id,
			mt.material_type_name,
			COUNT(mi.material_id) AS total_count,
			COALESCE(SUM(CASE WHEN mi.material_status = 0 THEN 1 ELSE 0 END), 0) AS available_count`).
		Joins("LEFT JOIN materials_info mi ON mi.material_type_id = mt.material_type_id").
		Group("mt.material_type_id, mt.material_type_name").
		Find(&stats).Error
	return stats, err
}
