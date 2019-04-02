package entity

import (
	. "gserver_service_demo/src/server/fight"
	. "gserver_service_demo/src/server/fight/cfg"

	. "github.com/gfandada/gserver/gameutil/entity"
)

type Tower struct {
	Entity
	tower  *Ai
	flag   int32
	space  SpaceId
	entity EntityId
}

func NewTower(flag int32, spaceid SpaceId, entityid EntityId) *Tower {
	return &Tower{
		tower:  NewTowerAi(),
		flag:   flag,
		space:  spaceid,
		entity: entityid,
	}
}

// 创建一个默认防御塔
// @params flag:类型 name:名称 spaceid:场景id
func NewDefaultTower(flag int32, name string, spaceid SpaceId) *Entity {
	tower := NewEntity(flag, name, true, true)
	durability, defend := GetTowerDura(uint32(1))
	if durability > 0 {
		tower.Increase(HP, float32(durability))
		tower.Increase(HP_OLD, float32(durability))
	}
	if defend > 0 {
		tower.Increase(DEFENSE, float32(defend))
	}
	tower.BindIentity(NewTower(flag, spaceid, tower.Id))
	RegisterEntity(tower)
	return tower
}

/*********************************实现Ientity接口********************************/

func (e *Tower) OnInit() {
}

func (e *Tower) OnCreated() {
}

func (e *Tower) OnDestroy() {
}

func (e *Tower) OnMigrateOut() {
}

func (e *Tower) OnMigrateIn() {
}

func (e *Tower) OnRestored() {
}

func (e *Tower) OnEnterSpace() {
}

func (e *Tower) OnLeaveSpace(space *Space) {
}

func (e *Tower) IsPersistent() bool {
	return false
}

func (e *Tower) Flag() int32 {
	return e.flag
}
