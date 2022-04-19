package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"comment:用户名" comment:"用户名" json:"username" form:"username" binding:"required,min=6"`
	Nickname string `gorm:"comment:昵称" comment:"昵称" json:"nickname" form:"nickname" binding:"required"`
	Age      uint8  `gorm:"comment:年龄" comment:"年龄" json:"age" form:"age"`
	Password string `gorm:"comment:密码" comment:"密码" json:"password" form:"password" binding:"required,min=6"`
	Gender   bool   `gorm:"comment:性别" comment:"性别" json:"gender" form:"gender"`
}
