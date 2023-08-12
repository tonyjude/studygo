package process2

import "fmt"

var (
	userMgr *UserMgr
)

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

//用户消息
type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//添加在线用户
func (_this *UserMgr) AddOnlineUser(up *UserProcess) {
	_this.onlineUsers[up.UserId] = up
}

//删除
func (_this *UserMgr) DelOnlineUser(userId int) {
	delete(_this.onlineUsers, userId)
}

//返回当前所有在线用户
func (_this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return _this.onlineUsers
}

func (_this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := _this.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%d不存在", userId)
		return
	}
	return
}
