package handler

import (
	"context"
	"thresher/srv/user-srv/util"
	"time"
)

var (
	UserList map[string]*User
)

type User struct {
	Id       int64
	Username string
	Password string
	Profile  Profile
}

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
}
// 定义请求结构体
type UserRequest struct {
	UserId int64 	`form:"user_id" json:"user_id"`
	UserName string `form:"user_name" json:"user_name"`
	PassWord string `form:"pass_word" json:"pass_word"`
}

//Response 结构体
type Response struct {
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}

func init() {
	UserList = make(map[string]*User)
	u := User{123, "admin", "admin", Profile{"male", 20, "Singapore", "wule61@live.com"}}
	UserList["123"] = &u
}

// 登录后根据用户名和密码生成jwt
func (u *User) Login(ctx context.Context,args UserRequest,reply *Response) (err error) {
	for _, v := range UserList {
		if v.Username == args.UserName && v.Password == args.PassWord{
			et := &util.EasyToken{
				Username: args.UserName,
				Uid: v.Id,
				Expires: time.Now().Unix() + 3600,
			}
			token,err := et.GetToken()
			if err != nil {
				reply.Errmsg = "生成jwt错误"
				return  err
			}
			reply.Errcode= 200
			reply.Data = map[string]interface{}{
				"jwt":token,
			}
		}
	}
	return nil
}