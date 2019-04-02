package entity

import (
	"time"

	. "gserver_service_demo/src/server/fight"
	. "gserver_service_demo/src/server/fight/cfg"

	. "github.com/gfandada/gserver/gameutil/entity"
	. "github.com/gfandada/gserver/gameutil/fight"
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
	. "github.com/magicsea/behavior3go/core"
	"github.com/magicsea/behavior3go/loader"
)

type SoldierInstanceCondition struct {
	Condition
	instance int
}

func (this *SoldierInstanceCondition) Initialize(setting *BTNodeCfg) {
	this.Condition.Initialize(setting)
	this.instance = setting.GetPropertyAsInt("instance")
}

func (this *SoldierInstanceCondition) OnTick(tick *Tick) b3.Status {
	entityid := tick.Blackboard.Get("entity", "entity", "entity").(EntityId)
	fightId := tick.Blackboard.Get("entity", "fightScheduler", "entity").(FightId)
	radius := tick.Blackboard.Get("entity", "radius", "entity").(uint32)
	attack := tick.Blackboard.Get("entity", "attack", "entity").(uint32)
	entity := GetEntity(entityid)
	if entity == nil {
		return b3.FAILURE
	}
	enemy, err := CallFightScheduler(fightId, GET_AOIS, []interface{}{entityid, int(radius)})
	if err != nil || len(enemy) == 0 {
		return b3.FAILURE
	}
	CastFightDamageCalc(fightId, AUTO_ATTACK, []interface{}{string(entityid),
		string(enemy[0].(EntityId)), int32(attack)})
	return b3.SUCCESS
}

type SoldierAttackAction struct {
	Action
	attack int
}

func (this *SoldierAttackAction) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
	this.attack = setting.GetPropertyAsInt("attack")
}

func (this *SoldierAttackAction) OnTick(tick *Tick) b3.Status {
	return b3.SUCCESS
}

func NewSoldierAi() *Ai {
	maps := b3.NewRegisterStructMaps()
	maps.Register("SoldierInstanceCondition", new(SoldierInstanceCondition))
	maps.Register("SoldierAttackAction", new(SoldierAttackAction))
	return &Ai{
		Tree:  loader.CreateBevTreeFromConfig(_aiManager.get("soldier.json"), maps),
		Black: NewBlackboard(),
	}
}

// 添加小兵ai
func AddSoldierAi(timer *FightTimer, fightId FightId, soldier *Entity) {
	attack, radius, duration := GetSoldierAttack(uint32(1))
	timer.AddRepeatJob(soldier.Id, time.Duration(duration)*time.Millisecond, 0, func(args []interface{}) {
		ai := soldier.I.(*Soldier).soldier
		ai.Black.Set("entity", soldier.Id, "entity", "entity")
		ai.Black.Set("entity", fightId, "fightScheduler", "entity")
		ai.Black.Set("entity", attack, "attack", "entity")
		ai.Black.Set("entity", radius, "radius", "entity")
		ai.Black.Set("entity", duration, "duration", "entity")
		ai.Tree.Tick(0, ai.Black)
	}, []interface{}{soldier})
}
