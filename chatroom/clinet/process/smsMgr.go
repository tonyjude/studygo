package process

import (
	"fmt"

	"../../message"
)

func getGroupMes(smsMes *message.SmsMes) {
	info := fmt.Sprintf("接收到群消息，用户ID:\t%d,内容是:\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println()

}
