package dao

import (
	"materials-service/internal/model"
	"fmt"
)



// CreateMaterialType 创建种类
func CreateMaterialType(materialType *model.MaterialType) error {
	result := model.DB.Create(materialType)
	if result.Error != nil {
		return fmt.Errorf("创建失败: %v", result.Error)
	}
	return nil
}

// GetMaterialTypeByID 根据ID查询种类
func GetMaterialTypeByID(id uint) (*model.MaterialType, error) {
	var materialType model.MaterialType
	result := model.DB.First(&materialType, id)
	if result.Error != nil {
		return nil, fmt.Errorf("查询失败: %v", result.Error)
	}
	return &materialType, nil
}

// GetAllMaterialTypes 查询所有种类
func GetAllMaterialTypes() ([]model.MaterialType, error) {
	var materialTypes []model.MaterialType
	result := model.DB.Find(&materialTypes)
	if result.Error != nil {
		return nil, fmt.Errorf("查询失败: %v", result.Error)
	}
	return materialTypes, nil
}

// UpdateMaterialType 更新种类
func UpdateMaterialType(materialType *model.MaterialType) error {
	result := model.DB.Save(materialType)
	if result.Error != nil {
		return fmt.Errorf("更新失败: %v", result.Error)
	}
	return nil
}

// DeleteMaterialType 删除种类（按ID）
func DeleteMaterialType(id uint) error {
	result := model.DB.Delete(&model.MaterialType{}, id)
	if result.Error != nil {
		return fmt.Errorf("删除失败: %v", result.Error)
	}
	return nil
}
