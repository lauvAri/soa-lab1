package main

import (
    "materials-service/internal/model"
    "materials-service/internal/config"
    //"materials-service/internal/dao"
    "materials-service/routers"
    //"net/http"
    //"github.com/gin-gonic/gin"
    "fmt"
)

// 定义物资结构体
type Material struct {
    ID   string `json:"id"`
    Name string `json:"name"`
    Type string `json:"type"`
}

// 模拟数据库
var materials = []Material{
    {ID: "1", Name: "显微镜", Type: "光学仪器"},
    {ID: "2", Name: "示波器", Type: "电子仪器"},
}

//测试数据
var testType = []model.MaterialType{
    {ID: 1, Name: "Mac"},
}

func main() {
    config.LoadEnv()
    // 初始化数据库
	if err := model.InitDB(); err != nil {
		fmt.Println("数据库初始化失败: %v", err)
	}
	defer func() {
		if db, _ := model.DB.DB(); db != nil {
			_ = db.Close()
		}
	}()

    //dao.CreateMaterialType(&testType[0])

    r := routers.SetupRouter(model.DB)

    // 运行在 8082 端口
    r.Run(":8082")
}