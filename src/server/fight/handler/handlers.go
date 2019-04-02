package handler

import (
	. "gserver_service_demo/src/server/fight"

	. "github.com/gfandada/gserver/gameutil/fight"
)

func NewFightHandlers() {
	// for fightScheduler
	RegisterHandler(INIT_SCHEDULER, initFight)
	RegisterHandler(CLOSE_SCHEDULER, closeFight)
	RegisterHandler(TIMER_SCHEDULER, timerFight)
	RegisterHandler(FLUSH_SOLDIER, flushSoldier)
	RegisterHandler(JOIN_FIGHT, joinFight)
	RegisterHandler(GET_FIGHT, getFight)
	RegisterHandler(LEAVE_FIGHT, leaveFight)
	RegisterHandler(MOVE_FIGHT, moveFight)
	RegisterHandler(GET_AOIS, getAois)
	RegisterHandler(ATTACK_ADAP, attackAdap)
	RegisterHandler(ADD_EQUIP, addEquip)
	RegisterHandler(RM_EQUIP, rmEquip)
	RegisterHandler(ROADCAST_TMP, broadcast)
	RegisterHandler(ATTACK_KILL, attackKill)
	// for fightdamage
	RegisterHandler(AUTO_ATTACK, attack)
	// for fightaward
	RegisterHandler(KILL, kill)
	RegisterHandler(ASSISTKILL, assistingKill)
	// TODO for fightpost 暂时未使用独立的POST服务
	RegisterHandler(ENTER_POST, enterPost)
	RegisterHandler(AOI_POST, aoiPost)
	RegisterHandler(SPWAN_ENTITY_POST, entitySpwanPost)
	RegisterHandler(DEATH_ENTITY_POST, entityDeathPost)
	RegisterHandler(GAME_OVER_POST, gameOverPost)
	RegisterHandler(BROADCAST, broadcastPost)
}
