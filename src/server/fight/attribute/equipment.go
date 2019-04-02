package attribute

import (
	. "gserver_service_demo/src/server/fight"
	. "gserver_service_demo/src/server/fight/cfg"
	. "gserver_service_demo/src/server/fight/entity"

	. "github.com/gfandada/gserver/gameutil/entity"
	. "github.com/gfandada/gserver/gameutil/fight"
)

func AddEquip(equipId uint32, fightId FightId, entityId EntityId, timer *FightTimer) {
	entitiy := GetEntity(entityId)
	if entitiy == nil {
		return
	}
	durability, defend := GetEquipDura(equipId)
	if durability > 0 {
		entitiy.Increase(HP, float32(durability))
		entitiy.Increase(HP_OLD, float32(durability))
	}
	if defend > 0 {
		entitiy.Increase(DEFENSE, float32(defend))
	}
	attack, radius, duration := GetEquipAttack(equipId)
	if duration > 0 {
		AddShipAi(timer, fightId, entityId, equipId, attack, radius, duration)
	}
}
