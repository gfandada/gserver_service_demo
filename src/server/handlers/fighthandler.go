package handlers

import (
	. "gserver_service_demo/src/server/fight"
	Msg "gserver_service_demo/src/server/protomsg"

	. "github.com/gfandada/gserver/gameutil/entity"
	. "github.com/gfandada/gserver/gameutil/fight"
	. "github.com/gfandada/gserver/goroutine"
	"github.com/gfandada/gserver/logger"
	"github.com/gfandada/gserver/network"
	Service "github.com/gfandada/gserver/services/service"
)

// 创建战斗
func createFightHandler(args []interface{}) []interface{} {
	data := args[0].(*network.RawMessage).MsgData
	fightId := data.(*Msg.CreateFightReq).Fightid
	err := CreateFight(FightId(*fightId), 1, nil)
	if err != nil {
		return nil
	}
	return nil
}

// 加入战斗
func joinFightHandler(args []interface{}) []interface{} {
	data := args[0].(*network.RawMessage).MsgData
	fightId := data.(*Msg.JoinFightReq).Fightid
	usertype := data.(*Msg.JoinFightReq).Type
	sess := args[1].(*Service.Session)
	//sessFightId := sess.Get(SESS_FIGHTID)
	userId := args[1].(*Service.Session).UserId
	//	if sessFightId != nil {
	//		panic(int(Msg.GameConst_FIGHTID3V3ERROR))
	//	}
	fightIdType := FightId(*fightId)
	v := QueryByName(NewFightSchedulerAlias(fightIdType))
	if v == nil {
		panic(int(Msg.GameConst_FIGHT3V3ERROR))
	}
	// 存储战斗id
	sess.AddData(SESS_FIGHTID, fightIdType)
	join, err := CallFightScheduler(fightIdType, JOIN_FIGHT, []interface{}{userId, *usertype})
	logger.Debug("user %d JoinFight return %v", userId, join)
	if err != nil {
		panic(int(Msg.GameConst_JOIN3V3ERROR))
	}
	// 存储自己的实体id
	sess.AddData(SESS_ENTITYID, join[0].(EntityId))
	users, err := CallFightScheduler(fightIdType, GET_FIGHT, []interface{}{userId})
	logger.Debug("user %d GetFightUsers return %v", userId, users)
	if err != nil {
		panic(int(Msg.GameConst_GET3V3ERROR))
	}
	return []interface{}{network.RawMessage{
		MsgId: uint16(2103),
		MsgData: &Msg.JoinFightAck{
			Spawns: users[0].([]*Msg.SpawnState),
		}}}
}

// 移动
func moveFightHandler(args []interface{}) []interface{} {
	sess := args[1].(*Service.Session)
	fightId := sess.Get(SESS_FIGHTID).(FightId)
	if fightId == "" {
		panic(int(Msg.GameConst_FIGHT3V3ERROR))
	}
	entityid := sess.Get(SESS_ENTITYID)
	if entityid == nil || entityid == "" {
		panic(int(Msg.GameConst_ENTITYIDERROR))
	}
	data := args[0].(*network.RawMessage).MsgData
	pos := data.(*Msg.MoveReq).Pos
	v := data.(*Msg.MoveReq).V
	CastFightScheduler(fightId, MOVE_FIGHT, []interface{}{
		entityid.(EntityId), *pos.X, *pos.Y, *v.X, *v.Y,
	})
	return nil
}
