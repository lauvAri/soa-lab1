package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"materials-service/internal/model"
	"materials-service/internal/service"
)

type MaterialController struct {
	svc *service.MaterialService
}

func NewMaterialController() *MaterialController {
	return &MaterialController{svc: service.NewMaterialService()}
}

// MaterialListResponse 描述列表接口的响应结构，仅用于 Swagger 文档。
type MaterialListResponse struct {
	Items    []model.Material `json:"items"`
	Total    int64            `json:"total"`
	Page     int              `json:"page"`
	PageSize int              `json:"pageSize"`
	Ts       time.Time        `json:"ts"`
}

// RegisterMaterialRoutes 将物资相关路由注册到 gin.IRouter
func RegisterMaterialRoutes(r gin.IRouter) {
	c := NewMaterialController()
	g := r.Group("/materials")
	{
		// 兼容带/不带尾随斜杠
		g.GET("", c.list)
		g.GET("/", c.list)
		g.GET("/stats", c.stats)
		g.GET("/:id", c.get)
		g.POST("", c.create)
		g.POST("/", c.create)
		g.PUT("/:id", c.update)
		g.DELETE("/:id", c.remove)
	}
}

// create 创建物资
// @Summary 创建物资
// @Description 新增一条物资记录
// @Tags 物资管理
// @Accept json
// @Produce json
// @Param material body model.Material true "物资信息"
// @Success 201 {object} model.Material
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /materials [post]
func (mc *MaterialController) create(c *gin.Context) {
	var req model.Material
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 容错：时间字段可用字符串时解析（可选）
	// 这里保持简洁，假设客户端按 RFC3339 发送时间
	created, err := mc.svc.Create(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// get 获取单个物资
// @Summary 获取物资详情
// @Tags 物资管理
// @Produce json
// @Param id path int true "物资ID"
// @Success 200 {object} model.Material
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /materials/{id} [get]
func (mc *MaterialController) get(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	mat, err := mc.svc.Get(c.Request.Context(), id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mat)
}

// update 更新物资
// @Summary 更新物资
// @Tags 物资管理
// @Accept json
// @Produce json
// @Param id path int true "物资ID"
// @Param material body model.Material true "要更新的字段"
// @Success 200 {object} model.Material
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /materials/{id} [put]
func (mc *MaterialController) update(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req model.Material
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := mc.svc.Update(c.Request.Context(), id, &req)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// remove 删除物资
// @Summary 删除物资
// @Tags 物资管理
// @Produce json
// @Param id path int true "物资ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /materials/{id} [delete]
func (mc *MaterialController) remove(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := mc.svc.Delete(c.Request.Context(), id); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// list 列出物资（分页）
// @Summary 物资列表
// @Tags 物资管理
// @Produce json
// @Param page query int false "页码，默认 1"
// @Param pageSize query int false "每页数量，默认 20"
// @Success 200 {object} controller.MaterialListResponse
// @Failure 500 {object} map[string]string
// @Router /materials [get]
func (mc *MaterialController) list(c *gin.Context) {
	page := mustAtoiDefault(c.Query("page"), 1)
	pageSize := mustAtoiDefault(c.Query("pageSize"), 20)
	items, total, err := mc.svc.List(c.Request.Context(), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"items":    items,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
		"ts":       time.Now().UTC(),
	})
}

// stats 返回物资统计信息
// @Summary 物资统计
// @Tags 物资管理
// @Produce json
// @Success 200 {object} model.MaterialStats
// @Failure 500 {object} map[string]string
// @Router /materials/stats [get]
func (mc *MaterialController) stats(c *gin.Context) {
	stats, err := mc.svc.Stats(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stats)
}

// 工具函数
func parseID(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func mustAtoiDefault(s string, def int) int {
	if s == "" {
		return def
	}
	v, err := strconv.Atoi(s)
	if err != nil || v <= 0 {
		return def
	}
	return v
}
