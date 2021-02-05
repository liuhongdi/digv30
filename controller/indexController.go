package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv30/global"
	"time"
)

type IndexController struct{}
func NewIndexController() IndexController {
	return IndexController{}
}
//首页，返回一个成功的提示
func (g *IndexController) Index(c *gin.Context) {
	fmt.Println("controller:index: "+time.Now().String())
	result := global.NewResult(c)
	result.Success("success");
	return
}
