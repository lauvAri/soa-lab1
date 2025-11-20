package dao  // 假设这个文件在 `dao` 包下

import (
	"fmt"
	"materials-service/internal/model"  // 确保路径正确
	"gorm.io/gorm"
)

type MaterialTypeDAO struct {
	DB *gorm.DB  // 依赖 GORM 数据库连接
}

// CreateMaterialType 创建种类（绑定到 MaterialTypeDAO）
func (dao *MaterialTypeDAO) CreateMaterialType(materialType *model.MaterialType) error {
	result := dao.DB.Create(materialType)
	if result.Error != nil {
		return fmt.Errorf("创建失败: %v", result.Error)
	}
	return nil
}

// GetMaterialTypeByID 根据ID查询种类（绑定到 MaterialTypeDAO）
func (dao *MaterialTypeDAO) GetMaterialTypeByID(id uint) (*model.MaterialType, error) {
	var materialType model.MaterialType
	result := dao.DB.First(&materialType, id)
	if result.Error != nil {
		return nil, fmt.Errorf("查询失败: %v", result.Error)
	}
	return &materialType, nil
}

// GetAllMaterialTypes 查询所有种类（绑定到 MaterialTypeDAO）
func (dao *MaterialTypeDAO) GetAllMaterialTypes() ([]model.MaterialType, error) {
	var materialTypes []model.MaterialType
	result := dao.DB.Find(&materialTypes)
	if result.Error != nil {
		return nil, fmt.Errorf("查询失败: %v", result.Error)
	}
	return materialTypes, nil
}

// UpdateMaterialType 更新种类（绑定到 MaterialTypeDAO）
func (dao *MaterialTypeDAO) UpdateMaterialType(materialType *model.MaterialType) error {
	result := dao.DB.Save(materialType)
	if result.Error != nil {
		return fmt.Errorf("更新失败: %v", result.Error)
	}
	return nil
}

// DeleteMaterialType 删除种类（绑定到 MaterialTypeDAO）
func (dao *MaterialTypeDAO) DeleteMaterialType(id uint) error {
	result := dao.DB.Delete(&model.MaterialType{}, id)
	if result.Error != nil {
		return fmt.Errorf("删除失败: %v", result.Error)
	}
	return nil
}
