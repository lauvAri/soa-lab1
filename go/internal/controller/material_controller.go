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

// RegisterMaterialRoutes 将物资相关路由注册到 gin.IRouter
func RegisterMaterialRoutes(r gin.IRouter) {
	c := NewMaterialController()
	g := r.Group("/materials")
	{
		// 兼容带/不带尾随斜杠
		g.GET("", c.list)
		g.GET("/", c.list)
		g.GET("/:id", c.get)
		g.POST("", c.create)
		g.POST("/", c.create)
		g.PUT("/:id", c.update)
		g.DELETE("/:id", c.remove)
	}
}

// create 创建物资
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
