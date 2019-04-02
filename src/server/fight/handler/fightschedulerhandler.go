package handler

import (
	"strconv"
	"time"

	. "gserver_service_demo/src/server/fight"
	. "gserver_service_demo/src/server/fight/attribute"
	. "gserver_service_demo/src/server/fight/cfg"
	. "gserver_service_demo/src/server/fight/entity"
	. "gserver_service_demo/src/server/protomsg"

	. "github.com/gfandada/gserver/gameutil/entity"
	. "github.com/gfandada/gserver/gameutil/fight"
	"github.com/gfandada/gserver/loader"
	"github.com/gfandada/gserver/logger"
)

// 初始化
func initFight(inner, args []interface{}) []interface{} {
	_, _, fightId, space, _, _, towers, timer := ParseSchedulerInner(inner)
	spaceId := space.Id
	flushTower(spaceId, fightId, towers, timer)
	flushCrystal(spaceId, fightId, towers, timer)
	return nil
}

// 清理
func closeFight(inner, args []interface{}) []interface{} {
	return nil
}

// 帧函数
func timerFight(inner, args []interface{}) []interface{} {
	data, _, _, _, ships, soldiers, _, _ := ParseSchedulerInner(inner)
	time := time.Now().UnixNano() / 1e6
	// 按匀速推导
	for ship := range ships {
		pos := ship.GetPosition()
		timej := Coord(time-pos.TIME) / Coord(1e3)
		newpos := Vector3{
			X:    pos.X + pos.VX*timej,
			Y:    Coord(0),
			Z:    pos.Z + pos.VZ*timej,
			VX:   pos.VX,
			VZ:   pos.VZ,
			TIME: time,
		}
		ship.MoveSpace(newpos)
	}
	data.(*FightMap).updateIndex(soldiers, time)
	PostAoi(ships, soldiers)
	return nil
}

// 刷新防御塔
func flushTower(spaceId SpaceId, fightId FightId, towers map[*Entity]struct{},
	timer *FightTimer) {
	loader := new(loader.Loader)
	for i := 1; i <= 5; i++ {
		tower := NewDefaultTower(int32(ENTITY_BLUE_TOWER_FLAG), strconv.Itoa(i),
			spaceId)
		RegisterEntity(tower)
		x, _ := loader.GetFloat64(loader.Get("MapSite", uint32(i), "PositionX"))
		z, _ := loader.GetFloat64(loader.Get("MapSite", uint32(i), "PositionY"))
		tower.EnterSpace(spaceId, Vector3{
			X:    Coord(float32(x)),
			Y:    Coord(0),
			Z:    Coord(float32(z)),
			VX:   Coord(0),
			VZ:   Coord(0),
			TIME: 0,
		})
		towers[tower] = struct{}{}
		AddTowerAi(timer, fightId, tower)
	}
	for i := 6; i <= 10; i++ {
		tower := NewDefaultTower(int32(ENTITY_RED_TOWER_FLAG), strconv.Itoa(i),
			spaceId)
		RegisterEntity(tower)
		x, _ := loader.GetFloat64(loader.Get("MapSite", uint32(i), "PositionX"))
		z, _ := loader.GetFloat64(loader.Get("MapSite", uint32(i), "PositionY"))
		tower.EnterSpace(spaceId, Vector3{
			X:    Coord(float32(x)),
			Y:    Coord(0),
			Z:    Coord(float32(z)),
			VX:   Coord(0),
			VZ:   Coord(0),
			TIME: 0,
		})
		towers[tower] = struct{}{}
		AddTowerAi(timer, fightId, tower)
	}
}

// 刷新水晶
func flushCrystal(spaceId SpaceId, fightId FightId, crystals map[*Entity]struct{},
	timer *FightTimer) {
	loader := new(loader.Loader)
	for i := 1; i <= 2; i++ {
		crystal := NewDefaultCrystal(int32(i+6),
			"crystal"+strconv.Itoa(i+6)+":"+strconv.Itoa(i),
			spaceId)
		RegisterEntity(crystal)
		x, _ := loader.GetFloat64(loader.Get("MapSite", uint32(i+10), "PositionX"))
		z, _ := loader.GetFloat64(loader.Get("MapSite", uint32(i+10), "PositionY"))
		crystal.EnterSpace(spaceId, Vector3{
			X:    Coord(float32(x)),
			Y:    Coord(0),
			Z:    Coord(float32(z)),
			VX:   Coord(0),
			VZ:   Coord(0),
			TIME: 0,
		})
		crystals[crystal] = struct{}{}
		AddCrystalAi(timer, fightId, crystal)
	}
}

// 刷新小兵
func flushSoldier(inner, args []interface{}) []interface{} {
	data, _, fightId, space, ships, soldiers, _, timer := ParseSchedulerInner(inner)
	spaceId := space.Id
	newsoldiers := make(map[*Entity]struct{})
	for j := BLUE_SHANG; j <= BLUE_XIA; j++ {
		for i := 0; i < 5; i++ {
			soldier := NewDefaultSoldier(int32(ENTITY_BLUE_SOLDIER_FLAG),
				"soldier"+strconv.Itoa(int(ENTITY_BLUE_SOLDIER_FLAG))+":"+strconv.Itoa(i),
				spaceId)
			RegisterEntity(soldier)
			x, y := data.(*FightMap).getPoint(j, i, soldier.Id)
			soldier.EnterSpace(spaceId, Vector3{
				X:    Coord(x),
				Y:    Coord(j),
				Z:    Coord(y),
				VX:   Coord(1),
				VZ:   Coord(1),
				TIME: 0,
			})
			soldiers[soldier] = struct{}{}
			newsoldiers[soldier] = struct{}{}
			AddSoldierAi(timer, fightId, soldier)
		}
	}
	for j := RED_SHANG; j <= RED_XIA; j++ {
		for i := 0; i < 5; i++ {
			soldier := NewDefaultSoldier(int32(ENTITY_RED_SOLDIER_FLAG),
				"soldier"+strconv.Itoa(int(ENTITY_RED_SOLDIER_FLAG))+":"+strconv.Itoa(i),
				spaceId)
			RegisterEntity(soldier)
			x, y := data.(*FightMap).getPoint(j, i, soldier.Id)
			soldier.EnterSpace(spaceId, Vector3{
				X:    Coord(x),
				Y:    Coord(j),
				Z:    Coord(y),
				VX:   Coord(1),
				VZ:   Coord(1),
				TIME: 0,
			})
			soldiers[soldier] = struct{}{}
			newsoldiers[soldier] = struct{}{}
			AddSoldierAi(timer, fightId, soldier)
		}
	}
	PostEntitySpwan(ships, newsoldiers)
	return nil
}

// 加入战斗
func joinFight(inner, args []interface{}) []interface{} {
	_, flag, fightId, space, ships, _, towers, timer := ParseSchedulerInner(inner)
	spaceId := space.Id
	ship := NewDefaultShip(args[1].(int32),
		"ship"+strconv.Itoa(int(args[1].(int32)))+":"+strconv.Itoa(int(args[0].(int32))),
		spaceId)
	ship.BindGameClient(args[0].(int32))
	RegisterEntity(ship)
	switch args[1].(int32) {
	case int32(ENTITY_BLUE_SHIP_FLAG):
		ship.EnterSpace(spaceId, Vector3{
			X: Coord(5 + len(ships)),
			Y: Coord(0),
			Z: Coord(5 + len(ships)),
		})
	case int32(ENTITY_RED_SHIP_FLAG):
		ship.EnterSpace(spaceId, Vector3{
			X: Coord(195 - len(ships)),
			Y: Coord(0),
			Z: Coord(145 - len(ships)),
		})
	}
	ships[ship] = struct{}{}
	PostEnter(ship, ships)
	// AddShipAi(timer, fightId, ship)
	addEquipAuto(fightId, timer, ship.Id)
	if flag*2 == len(ships) {
		timer.AddRepeatJob(EntityId("flushsoldier"), time.Duration(time.Second)*20, 0,
			func(args []interface{}) {
				CastFightScheduler(fightId, FLUSH_SOLDIER, []interface{}{})
			}, []interface{}{})
		PostEntitySpwan(ships, towers)
	}
	return []interface{}{ship.Id}
}

// 获取战斗双方数据
func getFight(inner, args []interface{}) []interface{} {
	_, _, _, _, ships, _, _, _ := ParseSchedulerInner(inner)
	return []interface{}{NewSpawnState(ships)}
}

// 离开战斗
func leaveFight(inner, args []interface{}) []interface{} {
	_, _, fightId, _, ships, _, _, timer := ParseSchedulerInner(inner)
	entityid := args[0].(EntityId)
	if entity := GetEntity(entityid); entity != nil {
		timer.DelAiJob(entity.Id)
		entity.LeaveSpace()
		UnRegisterEntity(entityid)
		delete(ships, entity)
		if len(ships) <= 0 {
			DestroyFight(fightId)
		}
	}
	return nil
}

// 移动指令
func moveFight(inner, args []interface{}) []interface{} {
	entityid := args[0].(EntityId)
	pos := Vector3{
		X:    Coord(args[1].(float32)),
		Y:    Coord(0),
		Z:    Coord(args[2].(float32)),
		VX:   Coord(args[3].(float32)),
		VZ:   Coord(args[4].(float32)),
		TIME: time.Now().UnixNano() / 1e6,
	}
	if entity := GetEntity(entityid); entity != nil {
		entity.MoveSpace(pos)
	}
	return nil
}

// 获取aoi信息
func getAois(inner, args []interface{}) []interface{} {
	entityId := args[0].(EntityId)
	radius := args[1].(int)
	entitiy := GetEntity(entityId)
	if entitiy == nil {
		return []interface{}{}
	}
	enemy := attackJudge(entitiy, entitiy.Neighbors(), radius)
	if enemy == "" {
		return []interface{}{}
	}
	return []interface{}{enemy}
}

// 攻击适配
func attackAdap(inner, args []interface{}) []interface{} {
	_, _, _, _, ships, _, _, _ := ParseSchedulerInner(inner)
	attacked := EntityId(args[1].(string))
	damage := args[2].(int32)
	hp := args[3].(int32)
	attackEntity := GetEntity(EntityId(args[0].(string)))
	attackedEntity := GetEntity(attacked)
	if attackEntity == nil && attackedEntity == nil {
		return nil
	}
	PostBroadcast(ships, NewAttackPush(args[0].(string), args[1].(string),
		damage, hp))
	return nil
}

// 击杀
func attackKill(inner, args []interface{}) []interface{} {
	_, _, fightId, _, ships, soldiers, towers, timer := ParseSchedulerInner(inner)
	attacked := EntityId(args[1].(string))
	attackedEntity := GetEntity(attacked)
	if attackedEntity == nil {
		logger.Info("不应该为nil")
		return nil
	}
	if attackedEntity.Desc.Flag == int32(ENTITY_BLUE_SOLDIER_FLAG) ||
		attackedEntity.Desc.Flag == int32(ENTITY_RED_SOLDIER_FLAG) {
		timer.DelAiJob(attacked)
		delete(soldiers, attackedEntity)
		attackedEntity.LeaveSpace()
		UnRegisterEntity(attacked)
	} else if attackedEntity.Desc.Flag == int32(ENTITY_BLUE_SHIP_FLAG) {
		attackedEntity.MoveSpace(Vector3{
			X:    Coord(0),
			Y:    Coord(0),
			Z:    Coord(0),
			VX:   Coord(0),
			VZ:   Coord(0),
			TIME: time.Now().UnixNano() / 1e6,
		})
		// 恢复当前耐久
		attackedEntity.Increase(HP, attackedEntity.GetAttr(HP_OLD))
	} else if attackedEntity.Desc.Flag == int32(ENTITY_RED_SHIP_FLAG) {
		attackedEntity.MoveSpace(Vector3{
			X:    Coord(49),
			Y:    Coord(0),
			Z:    Coord(49),
			VX:   Coord(0),
			VZ:   Coord(0),
			TIME: time.Now().UnixNano() / 1e6,
		})
		// 恢复当前耐久
		attackedEntity.Increase(HP, attackedEntity.GetAttr(HP_OLD))
	} else if attackedEntity.Desc.Flag == int32(ENTITY_BLUE_TOWER_FLAG) ||
		attackedEntity.Desc.Flag == int32(ENTITY_RED_TOWER_FLAG) {
		timer.DelAiJob(attacked)
		delete(towers, attackedEntity)
		attackedEntity.LeaveSpace()
		UnRegisterEntity(attacked)
	} else if attackedEntity.Desc.Flag == int32(ENTITY_BLUE_CRYSTAL) ||
		attackedEntity.Desc.Flag == int32(ENTITY_RED_CRYSTAL) {
		DestroyFight(fightId)
		PostBroadcast(ships, args[0])
		return nil
	}
	PostBroadcast(ships, args[0])
	return nil
}

// 添加装备
func addEquip(inner, args []interface{}) []interface{} {
	_, _, fightId, _, _, _, _, timer := ParseSchedulerInner(inner)
	entityId := args[0].(EntityId)
	equipId := args[1].(uint32)
	AddEquip(equipId, fightId, entityId, timer)
	return nil
}

// 添加装备
// TODO 系统自动
func addEquipAuto(fightId FightId, timer *FightTimer, entityId EntityId) {
	equips := GetDefaultShipEquip(uint32(1))
	for _, v := range equips {
		AddEquip(v, fightId, entityId, timer)
	}
}

// 移除装备
func rmEquip(inner, args []interface{}) []interface{} {
	return nil
}

// 广播
func broadcast(inner, args []interface{}) []interface{} {
	_, _, _, _, ships, _, _, _ := ParseSchedulerInner(inner)
	PostBroadcast(ships, args[0])
	return nil
}
