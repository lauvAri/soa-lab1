package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"materials-service/internal/dao"
	"materials-service/internal/model"
)

type MaterialTypeController struct {
	dao *dao.MaterialTypeDAO
}

func NewMaterialTypeController(dao *dao.MaterialTypeDAO) *MaterialTypeController {
	return &MaterialTypeController{dao: dao}
}

// CreateMaterialType 创建材料类型
// @Summary 创建材料类型
// @Tags 材料类型管理
// @Accept json
// @Produce json
// @Param materialType body model.MaterialType true "材料类型信息"
// @Success 201 {object} model.MaterialType
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /material-types [post]
func (c *MaterialTypeController) CreateMaterialType(ctx *gin.Context) {
	var materialType model.MaterialType
	if err := ctx.ShouldBindJSON(&materialType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	if err := c.dao.CreateMaterialType(&materialType); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, materialType)
}

// GetMaterialType 获取单个材料类型
// @Summary 获取材料类型详情
// @Tags 材料类型管理
// @Produce json
// @Param id path int true "材料类型ID"
// @Success 200 {object} model.MaterialType
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /material-types/{id} [get]
func (c *MaterialTypeController) GetMaterialType(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID格式"})
		return
	}

	materialType, err := c.dao.GetMaterialTypeByID(uint(id))
	if err != nil {
		if err.Error() == "查询失败: record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "材料类型不存在"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, materialType)
}

// ListMaterialTypes 获取所有材料类型
// @Summary 获取材料类型列表
// @Tags 材料类型管理
// @Produce json
// @Success 200 {array} model.MaterialType
// @Failure 500 {object} map[string]string
// @Router /material-types [get]
func (c *MaterialTypeController) ListMaterialTypes(ctx *gin.Context) {
	materialTypes, err := c.dao.GetAllMaterialTypes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, materialTypes)
}

// UpdateMaterialType 更新材料类型
// @Summary 更新材料类型
// @Tags 材料类型管理
// @Accept json
// @Produce json
// @Param id path int true "材料类型ID"
// @Param materialType body model.MaterialType true "更新后的材料类型信息"
// @Success 200 {object} model.MaterialType
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /material-types/{id} [put]
func (c *MaterialTypeController) UpdateMaterialType(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID格式"})
		return
	}

	var materialType model.MaterialType
	if err := ctx.ShouldBindJSON(&materialType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 确保更新的是指定ID的记录
	materialType.ID = uint(id)

	if err := c.dao.UpdateMaterialType(&materialType); err != nil {
		if err.Error() == "更新失败: record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "材料类型不存在"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, materialType)
}

// DeleteMaterialType 删除材料类型
// @Summary 删除材料类型
// @Tags 材料类型管理
// @Produce json
// @Param id path int true "材料类型ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /material-types/{id} [delete]
func (c *MaterialTypeController) DeleteMaterialType(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID格式"})
		return
	}

	if err := c.dao.DeleteMaterialType(uint(id)); err != nil {
		if err.Error() == "删除失败: record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "材料类型不存在"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.Status(http.StatusNoContent)
}

// RegisterMaterialTypeRoutes 将材料类型路由注册到 gin.IRouter
func RegisterMaterialTypeRoutes(r gin.IRouter, dao *dao.MaterialTypeDAO) {
	c := NewMaterialTypeController(dao)
	g := r.Group("/materials/types")
	{
		g.POST("", c.CreateMaterialType)
		g.GET("", c.ListMaterialTypes)
		g.GET("/:id", c.GetMaterialType)
		g.PUT("/:id", c.UpdateMaterialType)
		g.DELETE("/:id", c.DeleteMaterialType)
	}
}
