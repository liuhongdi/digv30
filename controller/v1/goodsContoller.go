package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv30/global"
)

type GoodsController struct{}
func NewGoodsController() GoodsController {
	return GoodsController{}
}
// v1 商品详情
func (g *GoodsController) GoodsOne(c *gin.Context) {
	result := global.NewResult(c)
	result.Success("v1 one");
	return
}
// v1 商品列表
func (g *GoodsController) GoodsList(c *gin.Context) {
	result := global.NewResult(c)
	result.Success("v1 list");
	return
}
