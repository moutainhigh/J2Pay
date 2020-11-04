package request

type RoleAdd struct {
	Role
}

type RoleEdit struct {
	ID int `json:"id"`
	Role
}
type Role struct {
	Pid  string   `json:"pid" example:"1"`                               // 上级ID
	Name string   `json:"name" binding:"required,max=255" example:"程序员"` // 角色名
	Auth []string `json:"auth" binding:"required,unique,min=1" example:"1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19"`
}
