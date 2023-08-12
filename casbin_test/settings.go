package main

// 模拟用户
var Users = map[string]User{
	"1":{
		Name:    "zhangsan",
		Age:     "20",
		RoleKey: admin,
	},
	"2":{
		Name:    "lisi",
		Age:     "21",
		RoleKey: formalMember,
	},
	"3":{
		Name:    "wangwu",
		Age:     "23",
		RoleKey: businessAdmin,
	},

	"4":{
		Name:    "zhaoliu",
		Age:     "24",
		RoleKey: normalUser,
	},

	"5":{
		Name:    "tianqi",
		Age:     "25",
		RoleKey: formalMember,
	},
}

//模拟认证
func GetIdentity(c *gin.Context) *User {
	auth := c.Request.Header.Get("Authentication")
	user,ok := Users[auth]
	if ok {
		return &user
	}
	return nil
}

