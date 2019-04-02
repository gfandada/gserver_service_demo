package handler

import (
	"fmt"

	. "gserver_service_demo/src/server/fight"

	. "gserver_service_demo/src/server/protomsg"

	. "github.com/gfandada/gserver/gameutil/entity"
	. "github.com/gfandada/gserver/gameutil/fight"
	"github.com/gfandada/gserver/logger"
)

// 攻击
func attack(inner, args []interface{}) []interface{} {
	fightId := ParseDamageCalcInner(inner)
	attack := args[0].(string)
	attacked := args[1].(string)
	damage := args[2].(int32)
	attackedEntity := GetEntity(EntityId(attacked))
	if attackedEntity != nil {
		ret := calcDamage(attackedEntity, EntityId(attack), EntityId(attacked), damage)
		if ret == 0 {
			// 击杀
			CastFightAward(fightId, KILL, []interface{}{attack, attacked, damage})
		} else {
			// 统计
			if attackedEntity.Desc.Flag == int32(ENTITY_BLUE_SHIP_FLAG) ||
				attackedEntity.Desc.Flag == int32(ENTITY_RED_SHIP_FLAG) {
				CastFightAward(fightId, ASSISTKILL, []interface{}{attack, attacked})
			}
			// 攻击适配
			CastFightScheduler(fightId, ATTACK_ADAP,
				[]interface{}{attack, attacked,
					damage, ret})
		}
	}
	return nil
}

// 伤害计算
// 造成伤害=攻击力*(1-(防御力*0.03)/(防御力*0.03+1))
// 耐久值=耐久值-造成伤害
func calcDamage(attackedEntity *Entity, attack, attacked EntityId, damge int32) int32 {
	defend := attackedEntity.GetAttr(DEFENSE)
	hp := attackedEntity.GetAttr(HP)
	calcd := float32(damge) * (1 - (defend*0.03)/(defend*0.03+1))
	hpNew := hp - calcd
	logger.Info(formatPrint(attack, attacked, damge, hp, hpNew))
	return int32(attackedEntity.Decrease(HP, calcd))
}

func GetName(entityId EntityId) string {
	entity := GetEntity(entityId)
	if entity == nil {
		return "这是死人"
	}
	switch entity.Desc.Flag {
	case int32(ENTITY_BLUE_SOLDIER_FLAG):
		return "蓝色小兵"
	case int32(ENTITY_RED_SOLDIER_FLAG):
		return "红色小兵"
	case int32(ENTITY_BLUE_SHIP_FLAG):
		return "蓝色战船"
	case int32(ENTITY_RED_SHIP_FLAG):
		return "红色战船"
	case int32(ENTITY_BLUE_TOWER_FLAG):
		return "蓝色防塔"
	case int32(ENTITY_RED_TOWER_FLAG):
		return "红色防塔"
	case int32(ENTITY_BLUE_CRYSTAL):
		return "蓝色水晶"
	case int32(ENTITY_RED_CRYSTAL):
		return "红色水晶"
	}
	return ""
}

// 格式化打印
func formatPrint(attack, attacked EntityId, damge int32, hp, hpNew float32) string {
	return fmt.Sprintf("[%s:%s]受到[%s:%s]%d点伤害导致耐久值降低[%f=>%f]", GetName(attacked), string(attacked),
		GetName(attack), string(attack), damge, hp, hpNew)
}

// 检查是否具备攻击对象
func attackJudge(entity *Entity, neighbors EntitySet, instance int) EntityId {
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
				return enemy.Id
			}
		}
		return ""
	} else if typeMyself == int32(ENTITY_RED_SHIP_FLAG) {
		for enemy := range entitys {
			if enemy.Desc.Flag == int32(ENTITY_RED_SHIP_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_RED_SOLDIER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_RED_TOWER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_RED_CRYSTAL) {
				continue
			}
			if entity.DistanceTo(enemy) <= Coord(instance) {
				return enemy.Id
			}
		}
		return ""
	} else if typeMyself == int32(ENTITY_BLUE_SOLDIER_FLAG) {
		for enemy := range entitys {
			if enemy.Desc.Flag == int32(ENTITY_BLUE_SHIP_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_BLUE_SOLDIER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_BLUE_TOWER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_BLUE_CRYSTAL) {
				continue
			}
			if entity.DistanceTo(enemy) <= Coord(instance) {
				return enemy.Id
			}
		}
		return ""
	} else if typeMyself == int32(ENTITY_RED_SOLDIER_FLAG) {
		for enemy := range entitys {
			if enemy.Desc.Flag == int32(ENTITY_RED_SHIP_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_RED_SOLDIER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_RED_TOWER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_RED_CRYSTAL) {
				continue
			}
			if entity.DistanceTo(enemy) <= Coord(instance) {
				return enemy.Id
			}
		}
		return ""
	} else if typeMyself == int32(ENTITY_BLUE_TOWER_FLAG) {
		for enemy := range entitys {
			if enemy.Desc.Flag == int32(ENTITY_BLUE_SHIP_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_BLUE_SOLDIER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_BLUE_TOWER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_BLUE_CRYSTAL) {
				continue
			}
			if entity.DistanceTo(enemy) <= Coord(instance) {
				return enemy.Id
			}
		}
		return ""
	} else if typeMyself == int32(ENTITY_RED_TOWER_FLAG) {
		for enemy := range entitys {
			if enemy.Desc.Flag == int32(ENTITY_RED_SHIP_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_RED_SOLDIER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_RED_TOWER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_RED_CRYSTAL) {
				continue
			}
			if entity.DistanceTo(enemy) <= Coord(instance) {
				return enemy.Id
			}
		}
		return ""
	} else if typeMyself == int32(ENTITY_RED_CRYSTAL) {
		for enemy := range entitys {
			if enemy.Desc.Flag == int32(ENTITY_RED_SHIP_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_RED_SOLDIER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_RED_TOWER_FLAG) {
				continue
			}
			if entity.DistanceTo(enemy) <= Coord(instance) {
				return enemy.Id
			}
		}
		return ""
	} else if typeMyself == int32(ENTITY_BLUE_CRYSTAL) {
		for enemy := range entitys {
			if enemy.Desc.Flag == int32(ENTITY_BLUE_SHIP_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_BLUE_SOLDIER_FLAG) ||
				enemy.Desc.Flag == int32(ENTITY_BLUE_TOWER_FLAG) {
				continue
			}
			if entity.DistanceTo(enemy) <= Coord(instance) {
				return enemy.Id
			}
		}
		return ""
	}
	return ""
}
