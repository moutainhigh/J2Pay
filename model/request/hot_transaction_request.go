package request

//创建hot交易
type HotTxAdd struct {
	From    string  `json:"from" binding:"required,max=255" example:"0xabcd"` //打币地址
	To      string  `json:"to" binding:"required,max=255" example:"0xabcd"`   //收币地址
	Balance float64 `json:"balance" binding:"required" example:"20"`          //金额
	Type    int     `json:"type" binding:"required" example:"1"`              //1:代发,2:排程结账,3:手动结账

}
