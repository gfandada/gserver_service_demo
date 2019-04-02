package handlers

import (
	Fight "gserver_service_demo/src/server/fight"

	Entity "github.com/gfandada/gserver/gameutil/entity"
	. "github.com/gfandada/gserver/gameutil/fight"
	Service "github.com/gfandada/gserver/services/service"
)

func closeHandler(args []interface{}) []interface{} {
	sess := args[0].(*Service.Session)
	//	userId := sess.UserId
	fightId := sess.Get(SESS_FIGHTID)
	entity := sess.Get(SESS_ENTITYID)
	if fightId == nil || entity == nil {
		return nil
	}
	entityId := entity.(Entity.EntityId)
	CastFightScheduler(fightId.(FightId), Fight.LEAVE_FIGHT, []interface{}{entityId})
	return nil
}
