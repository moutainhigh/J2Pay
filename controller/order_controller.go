package controller

import (
	"github.com/gin-gonic/gin"
	"j2pay-server/model"
	"j2pay-server/model/request"
	"j2pay-server/pkg/util"
	"j2pay-server/service"
	"net/http"
	"strconv"
)

// @Tags 订单管理
// @Summary 订单管理
// @Produce json
// @Router /orderIndex [get]
func OrderIndex(c *gin.Context) {
	c.HTML(200,"order.html" ,gin.H{
		"code": http.StatusOK,
	})
}

// @Tags 订单管理
// @Summary 订单列表
// @Produce json
// @Param status  query int false "状态 -1：收款中，1：已完成，2：异常，3：退款等待中，4：退款中，5：退款失败，6：已退款，7：已过期"
// @Param orderCode    query string false "商户订单编号"
// @Param userId    query int true "商户id"
// @Param txid    query string false "交易哈希"
// @Param chargeAddress    query string false "收款地址"
// @Param from_date  query string false "起"
// @Param to_date    query string false "至"
// @Param page query int false "页码"
// @Param pageSize query int false "每页显示多少条"
// @Success 200 {object} response.OrderPage
// @Router /order [get]
func OrderList(c *gin.Context) {
	response := util.Response{c}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	chargeAddress := c.Query("chargeAddress")
	orderCode := c.Query("orderCode")
	txid := c.Query("txid")
	FromDate := c.Query("fromDate")
	ToDate := c.Query("toDate")
	status, _ := strconv.Atoi(c.Query("status"))
	userId, _ := strconv.Atoi(c.Query("userId"))
	res, err := service.OrderList(FromDate,ToDate,status,chargeAddress,txid,orderCode,userId,page,pageSize)
	if err != nil {
		response.SetOtherError(err)
		return
	}
	response.SuccessData(res)
}

// @Tags 订单管理
// @Summary 订单详情
// @Produce json
// @Param id path uint true "ID"
// @Success 200 {object} response.RealOrderList
// @Router /order/{id} [get]
func OrderDetail(c *gin.Context) {
	response := util.Response{c}
	id, _ := strconv.Atoi(c.Param("id"))
	detail, err := service.OrderDetail(uint(id))
	if err != nil {
		response.SetOtherError(err)
		return
	}
	response.SuccessData(detail)
}

// @Tags 订单管理
// @Summary 新增订单(充币)
// @Produce json
// @Param body body request.OrderAdd true "新增订单"
// @Success 200
// @Router /order [post]
func OrderAdd(c *gin.Context) {
	response := util.Response{c}
	var order request.OrderAdd
	if err := c.ShouldBind(&order); err != nil {
		response.SetValidateError(err)
		return
	}
	err,orderAddr := service.OrderAdd(order)
	if err != nil {
		response.SetValidateError(err)
	}
	response.Send(200,"成功",orderAddr)

}

// @Tags 订单管理
// @Summary 修改订单
// @Produce json
// @Param id path int true "ID"
// @Param body body request.OrderEdit true "实收订单"
// @Router /order/{id} [put]
func OrderEdit(c *gin.Context) {
	response := util.Response{c}
	var order request.OrderEdit
	order.ID, _ = strconv.Atoi(c.Param("id"))
	if err := c.ShouldBind(&order); err != nil {
		response.SetValidateError(err)
		return
	}
	if err := service.OrderEdit(order); err != nil {
		response.SetOtherError(err)
		return
	}
	response.SuccessMsg("成功")
}

// @Tags 订单管理
// @Summary 通知
// @Produce json
// @Param token  query string true "token"
// @Success 302
// @Router /orderNotify [post]
func OrderNotify(c *gin.Context)  {
	response := util.Response{c}
	token := c.Query("token")
	adminUser ,_:= model.GetUserByWhere("token =?", token)
	c.Redirect(302, adminUser.ReturnUrl)
	c.Abort()
	response.SuccessMsg("成功")
}


