package entity

import (
	. "gserver_service_demo/src/server/fight"
	. "gserver_service_demo/src/server/fight/cfg"

	. "github.com/gfandada/gserver/gameutil/entity"
)

type Soldier struct {
	Entity
	soldier *Ai
	flag    int32
	space   SpaceId
	entity  EntityId
}

func NewSoldier(flag int32, spaceid SpaceId, entityid EntityId) *Soldier {
	return &Soldier{
		soldier: NewSoldierAi(),
		flag:    flag,
		space:   spaceid,
		entity:  entityid,
	}
}

// 创建一个默认小兵
// @params flag:类型 name:名称 spaceid:场景id
func NewDefaultSoldier(flag int32, name string, spaceid SpaceId) *Entity {
	soldier := NewEntity(flag, name, true, true)
	durability, defend := GetSoldierDura(uint32(1))
	if durability > 0 {
		soldier.Increase(HP, float32(durability))
		soldier.Increase(HP_OLD, float32(durability))
	}
	if defend > 0 {
		soldier.Increase(DEFENSE, float32(defend))
	}
	soldier.BindIentity(NewSoldier(flag, spaceid, soldier.Id))
	RegisterEntity(soldier)
	return soldier
}

/*********************************实现Ientity接口********************************/

func (e *Soldier) OnInit() {
}

func (e *Soldier) OnCreated() {
}

func (e *Soldier) OnDestroy() {
}

func (e *Soldier) OnMigrateOut() {
}

func (e *Soldier) OnMigrateIn() {
}

func (e *Soldier) OnRestored() {
}

func (e *Soldier) OnEnterSpace() {
}

func (e *Soldier) OnLeaveSpace(space *Space) {
}

func (e *Soldier) IsPersistent() bool {
	return false
}

func (e *Soldier) Flag() int32 {
	return e.flag
}
