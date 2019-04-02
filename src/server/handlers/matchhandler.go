package handlers

import (
	"gserver_service_demo/src/server/match"
	Msg "gserver_service_demo/src/server/protomsg"

	"github.com/gfandada/gserver/network"
	Service "github.com/gfandada/gserver/services/service"
)

// 请求匹配
// FIXME 暂未区分战斗类型
func matchReqHandler(args []interface{}) []interface{} {
	userId := args[1].(*Service.Session).UserId
	data := args[0].(*network.RawMessage).MsgData
	fightType := data.(*Msg.MatchReq).GetFighttype()
	match.Match(userId, int(fightType))
	return nil
}
