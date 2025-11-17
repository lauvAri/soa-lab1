package routers

import (
	"materials-service/internal/controller"
	"materials-service/internal/dao"
	//"materials-service/internal/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// 初始化DAO
	mtDao := &dao.MaterialTypeDAO{DB: db}

	// 统一由 controller 暴露的注册函数来绑定路由
	api := r.Group("/api/v1")
	controller.RegisterMaterialRoutes(api)
	controller.RegisterMaterialTypeRoutes(api, mtDao)

	return r
}
