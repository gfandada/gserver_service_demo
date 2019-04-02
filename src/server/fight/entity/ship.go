package entity

import (
	. "gserver_service_demo/src/server/fight"
	. "gserver_service_demo/src/server/fight/cfg"
	. "gserver_service_demo/src/server/protomsg"

	. "github.com/gfandada/gserver/gameutil/entity"
)

type Ship struct {
	Entity
	ship   *Ai
	flag   int32
	space  SpaceId
	entity EntityId
}

func NewShip(flag int32, spaceid SpaceId, entityid EntityId) *Ship {
	return &Ship{
		//ship:   NewShipAi(),
		flag:   flag,
		space:  spaceid,
		entity: entityid,
	}
}

// 创建一个默认战船
// @params flag:类型 name:名称 spaceid:场景id
func NewDefaultShip(flag int32, name string, spaceid SpaceId) *Entity {
	ship := NewEntity(flag, name, true, true)
	durability, defend := GetDefaultShip(uint32(1))
	if durability > 0 {
		ship.Increase(HP, float32(durability))
		ship.Increase(HP_OLD, float32(durability))
	}
	if defend > 0 {
		ship.Increase(DEFENSE, float32(defend))
	}
	ship.BindIentity(NewShip(flag, spaceid, ship.Id))
	RegisterEntity(ship)
	return ship
}

// 攻击判断
func AttackJudge(entity *Entity, neighbors EntitySet, instance int) *Entity {
	entitys := neighbors
	typeMyself := entity.Desc.Flag
	if typeMyself == int32(ENTITY_BLUE_SHIP_FLAG) {
		for enemy := range entitys {
			if enemy.Desc.Flag == int32(ENTITY_BLUE_SHIP_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_BLUE_SOLDIER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_BLUE_TOWER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_BLUE_CRYSTAL) {
				continue
			}
			if entity.DistanceTo(enemy) <= Coord(instance) {
				return enemy
			}
		}
		return nil
	} else if typeMyself == int32(ENTITY_RED_SHIP_FLAG) {
		for enemy := range entitys {
			if enemy.Desc.Flag == int32(ENTITY_RED_SHIP_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_RED_SOLDIER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_RED_TOWER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_RED_CRYSTAL) {
				continue
			}
			if entity.DistanceTo(enemy) <= Coord(instance) {
				return enemy
			}
		}
		return nil
	} else if typeMyself == int32(ENTITY_BLUE_SOLDIER_FLAG) {
		for enemy := range entitys {
			if enemy.Desc.Flag == int32(ENTITY_BLUE_SHIP_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_BLUE_SOLDIER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_BLUE_TOWER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_BLUE_CRYSTAL) {
				continue
			}
			if entity.DistanceTo(enemy) <= Coord(instance) {
				return enemy
			}
		}
		return nil
	} else if typeMyself == int32(ENTITY_RED_SOLDIER_FLAG) {
		for enemy := range entitys {
			if enemy.Desc.Flag == int32(ENTITY_RED_SHIP_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_RED_SOLDIER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_RED_TOWER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_RED_CRYSTAL) {
				continue
			}
			if entity.DistanceTo(enemy) <= Coord(instance) {
				return enemy
			}
		}
		return nil
	} else if typeMyself == int32(ENTITY_BLUE_TOWER_FLAG) {
		for enemy := range entitys {
			if enemy.Desc.Flag == int32(ENTITY_BLUE_SHIP_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_BLUE_SOLDIER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_BLUE_TOWER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_BLUE_CRYSTAL) {
				continue
			}
			if entity.DistanceTo(enemy) <= Coord(instance) {
				return enemy
			}
		}
		return nil
	} else if typeMyself == int32(ENTITY_RED_TOWER_FLAG) {
		for enemy := range entitys {
			if enemy.Desc.Flag == int32(ENTITY_RED_SHIP_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_RED_SOLDIER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_RED_TOWER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_RED_CRYSTAL) {
				continue
			}
			if entity.DistanceTo(enemy) <= Coord(instance) {
				return enemy
			}
		}
		return nil
	} else if typeMyself == int32(ENTITY_RED_CRYSTAL) {
		for enemy := range entitys {
			if enemy.Desc.Flag == int32(ENTITY_RED_SHIP_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_RED_SOLDIER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_RED_TOWER_FLAG) {
				continue
			}
			if entity.DistanceTo(enemy) <= Coord(instance) {
				return enemy
			}
		}
		return nil
	} else if typeMyself == int32(ENTITY_BLUE_CRYSTAL) {
		for enemy := range entitys {
			if enemy.Desc.Flag == int32(ENTITY_BLUE_SHIP_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_BLUE_SOLDIER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_BLUE_TOWER_FLAG) {
				continue
			}
			if entity.DistanceTo(enemy) <= Coord(instance) {
				return enemy
			}
		}
		return nil
	}
	return nil
}

/*********************************实现Ientity接口********************************/

func (e *Ship) OnInit() {
}

func (e *Ship) OnCreated() {
}

func (e *Ship) OnDestroy() {
}

func (e *Ship) OnMigrateOut() {
}

func (e *Ship) OnMigrateIn() {
}

func (e *Ship) OnRestored() {
}

func (e *Ship) OnEnterSpace() {
}

func (e *Ship) OnLeaveSpace(space *Space) {
}

func (e *Ship) IsPersistent() bool {
	return false
}

func (e *Ship) Flag() int32 {
	return e.flag
}
