package model

import (
	"net"

	"../../message"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
