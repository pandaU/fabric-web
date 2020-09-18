package models

type UserInfo struct {
	Id int `json:"id"`
	Name string `json:"name" gorm:"column:truename"`
	Role string `json:"role"` //角色
	Phone string `json:"phone"`
}
func (user UserInfo) TableName() string {
	return "sys_user"
}
