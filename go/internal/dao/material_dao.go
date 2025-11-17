package dao

import (
	"context"
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
	return model.DB.WithContext(ctx).Delete(&model.Material{}, "material_id = ?", id).Error
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
