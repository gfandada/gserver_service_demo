package entity

import (
	. "gserver_service_demo/src/server/fight"
	. "gserver_service_demo/src/server/fight/cfg"

	. "github.com/gfandada/gserver/gameutil/entity"
)

type Crystal struct {
	Entity
	ai     *Ai
	flag   int32
	space  SpaceId
	entity EntityId
}

func NewCrystal(flag int32, spaceid SpaceId, entityid EntityId) *Crystal {
	return &Crystal{
		ai:     NewCrystalAi(),
		flag:   flag,
		space:  spaceid,
		entity: entityid,
	}
}

// 创建一个默认水晶
// @params flag:类型 name:名称 spaceid:场景id
func NewDefaultCrystal(flag int32, name string, spaceid SpaceId) *Entity {
	entity := NewEntity(flag, name, true, true)
	durability, defend := GetTowerDura(uint32(21))
	if durability > 0 {
		entity.Increase(HP, float32(durability))
		entity.Increase(HP_OLD, float32(durability))
	}
	if defend > 0 {
		entity.Increase(DEFENSE, float32(defend))
	}
	entity.BindIentity(NewCrystal(flag, spaceid, entity.Id))
	RegisterEntity(entity)
	return entity
}

/*********************************实现Ientity接口********************************/

func (c *Crystal) OnInit() {
}

func (c *Crystal) OnCreated() {
}

func (c *Crystal) OnDestroy() {
}

func (c *Crystal) OnMigrateOut() {
}

func (c *Crystal) OnMigrateIn() {
}

func (c *Crystal) OnRestored() {
}

func (c *Crystal) OnEnterSpace() {
}

func (c *Crystal) OnLeaveSpace(space *Space) {
}

func (c *Crystal) IsPersistent() bool {
	return false
}

func (c *Crystal) Flag() int32 {
	return c.flag
}
