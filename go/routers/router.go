package routers

import (
	"materials-service/internal/controller"
	"materials-service/internal/dao"
	"materials-service/internal/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// 初始化DAO
	mtDao := &dao.MaterialTypeDAO{DB: db}

	// 初始化Controller
	mtController := controller.NewMaterialTypeController(mtDao)

	// 路由组
	api := r.Group("/api/v1")
	{
		api.POST("/material-types", mtController.CreateMaterialType)
		api.GET("/material-types", mtController.ListMaterialTypes)
		api.GET("/material-types/:id", mtController.GetMaterialType)
		api.PUT("/material-types/:id", mtController.UpdateMaterialType)
		api.DELETE("/material-types/:id", mtController.DeleteMaterialType)
	}

	return r
}