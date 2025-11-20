package service

import (
	"context"
	"errors"
	"materials-service/internal/dao"
	"materials-service/internal/model"
)

type MaterialService struct{}

func NewMaterialService() *MaterialService { return &MaterialService{} }

func (s *MaterialService) Create(ctx context.Context, m *model.Material) (*model.Material, error) {
	if m == nil {
		return nil, errors.New("material is nil")
	}
	// 防止客户端注入主键
	m.MaterialID = 0
	if err := dao.CreateMaterial(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}

func (s *MaterialService) Get(ctx context.Context, id int64) (*model.Material, error) {
	return dao.GetMaterialByID(ctx, id)
}

func (s *MaterialService) Update(ctx context.Context, id int64, patch *model.Material) (*model.Material, error) {
	if patch == nil {
		return nil, errors.New("material is nil")
	}
	patch.MaterialID = id
	if err := dao.UpdateMaterial(ctx, patch); err != nil {
		return nil, err
	}
	// 返回最新数据
	return dao.GetMaterialByID(ctx, id)
}

func (s *MaterialService) Delete(ctx context.Context, id int64) error {
	return dao.DeleteMaterial(ctx, id)
}

func (s *MaterialService) List(ctx context.Context, page, pageSize int) ([]model.Material, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	return dao.ListMaterials(ctx, offset, pageSize)
}

// Stats 返回物资统计信息
func (s *MaterialService) Stats(ctx context.Context) (*model.MaterialStats, error) {
	statsByType, err := dao.GetMaterialStats(ctx)
	if err != nil {
		return nil, err
	}
	var availableTotal int64
	for _, item := range statsByType {
		availableTotal += item.AvailableCount
	}
	return &model.MaterialStats{
		ByType:         statsByType,
		AvailableTotal: availableTotal,
	}, nil
}
