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

	// 初始化Controller
	mtController := controller.NewMaterialTypeController(mtDao)

	// 路由组
	api := r.Group("/api/v1")
	{
		api.POST("/materials/types", mtController.CreateMaterialType)
		api.GET("/materials/types", mtController.ListMaterialTypes)
		api.GET("/materials/types/:id", mtController.GetMaterialType)
		api.PUT("/materials/types/:id", mtController.UpdateMaterialType)
		api.DELETE("/materials/types/:id", mtController.DeleteMaterialType)
	}

	return r
}