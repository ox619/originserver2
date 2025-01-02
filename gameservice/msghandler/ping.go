package msghandler

import (
	"originserver2/common/proto/msg"
	"originserver2/gameservice/player"
)

func ping(player *player.Player, msg *msg.MsgNil) {
	player.Ping()
}
