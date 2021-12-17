package service

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "watch/docs"
	"watch/handler"
)

func InitHttp() *gin.Engine {
	r := gin.New()
	// http://localhost:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/service/submitInfo", handler.SubmitInfo)
	r.POST("/service/addNode", handler.AddNode)

	return r
}
