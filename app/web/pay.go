package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zkeai/cusdt/app/log"
	"github.com/zkeai/cusdt/app/model"
	"net/url"
	"time"
)

func CheckoutCounter(ctx *gin.Context) {
	var tradeId = ctx.Param("trade_id")
	var order, ok = model.GetTradeOrder(tradeId)
	if !ok {
		ctx.String(200, "订单不存在")

		return
	}

	uri, err := url.ParseRequestURI(order.ReturnUrl)
	if err != nil {
		ctx.String(200, "同步地址错误")
		log.Error("同步地址解析错误", err.Error())

		return
	}

	ctx.HTML(200, order.TradeType+".html", gin.H{
		"http_host":  uri.Host,
		"trade_id":   tradeId,
		"amount":     order.Amount,
		"address":    order.Address,
		"expire":     int64(order.ExpiredAt.Sub(time.Now()).Seconds()),
		"return_url": order.ReturnUrl,
		"usdt_rate":  order.TradeRate,
	})
}

func CheckStatus(ctx *gin.Context) {
	var tradeId = ctx.Param("trade_id")
	var order, ok = model.GetTradeOrder(tradeId)
	if !ok {
		ctx.JSON(200, RespFailJson(fmt.Errorf("订单不存在")))

		return
	}

	var returnUrl string
	if order.Status == model.OrderStatusSuccess {

		returnUrl = order.ReturnUrl
	}

	// 返回响应数据
	ctx.JSON(200, gin.H{"trade_id": tradeId, "status": order.Status, "return_url": returnUrl})
}
