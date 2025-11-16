package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
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

func main() {
    r := gin.Default()

    // 获取所有物资
    r.GET("/materials", func(c *gin.Context) {
        c.JSON(http.StatusOK, materials)
    })

    // 添加物资 (简化版，未做持久化)
    r.POST("/materials", func(c *gin.Context) {
        var newMat Material
        if err := c.ShouldBindJSON(&newMat); err == nil {
            materials = append(materials, newMat)
            c.JSON(http.StatusCreated, newMat)
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
    })

    // 运行在 8082 端口
    r.Run(":8082")
}