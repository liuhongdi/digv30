package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv30/controller"
	v1 "github.com/liuhongdi/digv30/controller/v1"
	v2 "github.com/liuhongdi/digv30/controller/v2"
	"github.com/liuhongdi/digv30/global"
	"log"
	"runtime/debug"
)

func Router() *gin.Engine {
	router := gin.Default()
	//处理异常
	router.NoRoute(HandleNotFound)
	router.NoMethod(HandleNotFound)
	router.Use(Recover)
	// 路径映射
	//无group
	indexc:=controller.NewIndexController()
	router.GET("/index/index", indexc.Index);
    //v1 group
	apiv1 := router.Group("/v1")
	{
		goodsc:=v1.NewGoodsController()
		apiv1.GET("/goods/one", goodsc.GoodsOne)
		apiv1.GET("/goods/list", goodsc.GoodsList)
		//v1.POST("/read", readEndpoint)
	}
	//v2 group
	apiv2 := router.Group("/v2")
	{
		goodsc:=v2.NewGoodsController()
		apiv2.GET("/goods/one", goodsc.GoodsOne)
		apiv2.GET("/goods/list", goodsc.GoodsList)
	}
	return router
}

func HandleNotFound(c *gin.Context) {
	global.NewResult(c).Error(404,"资源未找到")
	return
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			global.NewResult(c).Error(500,"服务器内部错误")
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}