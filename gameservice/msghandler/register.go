package msghandler

import (
	"originserver2/common/proto/msg"
	"originserver2/gameservice/msgrouter"
)

func init() {
	msgrouter.RegMsgHandler(msg.MsgType_Ping, ping)
}
