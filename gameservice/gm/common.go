package gm

import (
	"originserver2/common/proto/msg"
	"originserver2/gameservice/player"
)

type Common struct {
}

// TestMsg 模拟客户端发送消息
func (c *Common) TestMsg(p *player.Player, arg []string) string {
	msgTpe := msg.MsgType(getArgInt(arg, 0))

	p.GetMsgReceiver().GmReceiver(p, msgTpe, []byte(getArgString(arg, 1)))
	return OK
}
