package model

import (
	"encoding/json"
	"fmt"

	"../../message"
	"github.com/garyburd/redigo/redis"
)

//完成对USER结构体的增删改查

var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

func (_this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	return
}

func (_this *UserDao) Login(userId int, userPwd string) (user *User, err error) {

	conn := _this.pool.Get()
	defer conn.Close()
	user, err = _this.getUserById(conn, userId)
	if err != nil {
		return
	}

	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return

}

func (_this *UserDao) Register(user *message.User) (err error) {

	conn := _this.pool.Get()
	defer conn.Close()
	_, err = _this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		return
	}

	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("用户注册失败 err=", err)
		return
	}

	return
}
