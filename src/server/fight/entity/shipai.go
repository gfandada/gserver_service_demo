package entity

import (
	"time"

	. "gserver_service_demo/src/server/fight"

	. "github.com/gfandada/gserver/gameutil/entity"
	. "github.com/gfandada/gserver/gameutil/fight"
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
	. "github.com/magicsea/behavior3go/core"
	"github.com/magicsea/behavior3go/loader"
)

type ShipInstanceCondition struct {
	Condition
	instance int
}

func (this *ShipInstanceCondition) Initialize(setting *BTNodeCfg) {
	this.Condition.Initialize(setting)
	this.instance = setting.GetPropertyAsInt("instance")
}

func (this *ShipInstanceCondition) OnTick(tick *Tick) b3.Status {
	entityid := tick.Blackboard.Get("entity", "entity", "entity").(EntityId)
	fightId := tick.Blackboard.Get("entity", "fightScheduler", "entity").(FightId)
	radius := tick.Blackboard.Get("entity", "radius", "entity").(uint32)
	attack := tick.Blackboard.Get("entity", "attack", "entity").(uint32)
	entity := GetEntity(entityid)
	if entity == nil {
		return b3.FAILURE
	}
	//neighbors, err := GetAois(NewFightSchedulerAlias(fightId), entityid)
	enemy, err := CallFightScheduler(fightId, GET_AOIS, []interface{}{entityid, int(radius)})
	if err != nil || len(enemy) == 0 {
		return b3.FAILURE
	}
	CastFightDamageCalc(fightId, AUTO_ATTACK, []interface{}{string(entityid),
		string(enemy[0].(EntityId)), int32(attack)})
	return b3.SUCCESS
}

type ShipAttackAction struct {
	Action
	attack int
}

func (this *ShipAttackAction) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
	this.attack = setting.GetPropertyAsInt("attack")
}

func (this *ShipAttackAction) OnTick(tick *Tick) b3.Status {
	return b3.SUCCESS
}

func NewShipAi() *Ai {
	maps := b3.NewRegisterStructMaps()
	maps.Register("ShipInstanceCondition", new(ShipInstanceCondition))
	maps.Register("ShipAttackAction", new(ShipAttackAction))
	return &Ai{
		Tree:  loader.CreateBevTreeFromConfig(_aiManager.get("ship.json"), maps),
		Black: NewBlackboard(),
	}
}

// 添加ai
// TODO jobid设计的有问题 暂不修改
func AddShipAi(timer *FightTimer, fightId FightId, shipId EntityId,
	equipId, attack, radius, duration uint32) {
	timer.AddRepeatJob(shipId, time.Duration(duration)*time.Millisecond, 0, func(args []interface{}) {
		ai := NewShipAi()
		ai.Black.Set("entity", shipId, "entity", "entity")
		ai.Black.Set("entity", fightId, "fightScheduler", "entity")
		ai.Black.Set("entity", equipId, "equipid", "entity")
		ai.Black.Set("entity", attack, "attack", "entity")
		ai.Black.Set("entity", radius, "radius", "entity")
		ai.Black.Set("entity", duration, "duration", "entity")
		ai.Tree.Tick(0, ai.Black)
	}, []interface{}{})
}
