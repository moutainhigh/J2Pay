package request

type MessageAdd struct {
	CommonMessage
}

type MessageEdit struct {
	ID int `json:"id" form:"id"`
	CommonMessage
}

type CommonMessage struct {
	Title     string `json:"title" binding:"required,max=255" example:"这是一个系统公告"`                   // 系统公告
	BeginTime int64  `json:"begin_time" binding:"required" example:"12431421234" form:"begin_time"` // 开始时间
	EndTime   int64  `json:"end_time" binding:"required" example:"124124124" form:"end_time"`       // 结束时间
	Users     []int  `json:"users" binding:"required,min=1" example:"1,2,3,4,5,6,7" form:"users"`   //发给用户id
}
