package v2

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv30/global"
)

type GoodsController struct{}
func NewGoodsController() GoodsController {
	return GoodsController{}
}
// v2 商品详情
func (g *GoodsController) GoodsOne(c *gin.Context) {
	result := global.NewResult(c)
	result.Success("v2 one");
	return
}
// v2 商品列表
func (g *GoodsController) GoodsList(c *gin.Context) {
	result := global.NewResult(c)
	result.Success("v2 list");
	return
}
