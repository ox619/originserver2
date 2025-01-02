package interfacedef

import (
	"originserver2/common/proto/msg"

	"google.golang.org/protobuf/proto"
)

type IMsgSender interface {
	SendToClient(clientId string, msgType msg.MsgType, message proto.Message) int
	CastToClient(clientIdList []string, msgType msg.MsgType, message proto.Message) int
	SendToPlayer(playerUserId string, msgType msg.MsgType, message proto.Message) int
	CastToPlayer(playerUserId []string, msgType msg.MsgType, message proto.Message) int
}
