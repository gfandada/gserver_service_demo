package handler

import (
	"fmt"

	. "gserver_service_demo/src/server/fight"
	. "gserver_service_demo/src/server/protomsg"

	. "github.com/gfandada/gserver/gameutil/entity"
	. "github.com/gfandada/gserver/gameutil/fight"
	"github.com/gfandada/gserver/logger"
)

// 击杀
func kill(inner, args []interface{}) []interface{} {
	fightId, kill, statistics := ParseAwardInner(inner)
	// 攻击方
	attack := EntityId(args[0].(string))
	// 被攻击方
	attacked := EntityId(args[1].(string))
	// 伤害
	damage := args[2].(int32)
	// 统计助攻
	list := func() (ret []string) {
		for _, v := range kill[attacked] {
			ret = append(ret, string(v))
			if entity := GetEntity(v); entity != nil {
				if entity.Desc.Flag == int32(ENTITY_BLUE_SHIP_FLAG) ||
					entity.Desc.Flag == int32(ENTITY_RED_SHIP_FLAG) {
					if _, ok := statistics[v]; !ok {
						statistics[v] = []int{0, 0, 0}
					} else {
						statistics[v] = []int{
							statistics[v][0],
							statistics[v][1],
							statistics[v][2] + 1,
						}
					}
				}
			}
		}
		return
	}()
	attackedEntity := GetEntity(attacked)
	attackEntity := GetEntity(attack)
	// 无效
	if attackedEntity == nil || attackEntity == nil {
		return nil
	}
	// 水晶boom
	if attackedEntity.Desc.Flag == int32(ENTITY_BLUE_CRYSTAL) {
		CastFightScheduler(fightId, ATTACK_KILL, []interface{}{
			NewGameOverPush(int32(ENTITY_RED_CRYSTAL), statistics),
			args[1].(string),
		})
	} else if attackedEntity.Desc.Flag == int32(ENTITY_RED_CRYSTAL) {
		CastFightScheduler(fightId, ATTACK_KILL, []interface{}{
			NewGameOverPush(int32(ENTITY_BLUE_CRYSTAL), statistics),
			args[1].(string),
		})
	} else {
		CastFightScheduler(fightId, ATTACK_KILL, []interface{}{
			NewSpwanDeadPush(args[1].(string), args[0].(string), list, damage),
			args[1].(string),
		})
	}
	// 统计击杀和死亡
	if _, ok := statistics[attacked]; !ok {
		statistics[attacked] = []int{0, 0, 0}
	}
	statistics[attacked] = []int{
		statistics[attacked][0],
		statistics[attacked][1] + 1,
		statistics[attacked][2],
	}
	if attackEntity.Desc.Flag == int32(ENTITY_BLUE_SHIP_FLAG) ||
		attackEntity.Desc.Flag == int32(ENTITY_RED_SHIP_FLAG) {
		if _, ok := statistics[attack]; !ok {
			statistics[attack] = []int{0, 0, 0}
		}
		statistics[attack] = []int{
			statistics[attack][0] + 1,
			statistics[attack][1],
			statistics[attack][2],
		}
	}
	logger.Info(formatPrintAward(attack, attacked, list, statistics[attacked], statistics[attack]))
	// 清空本次击杀回放
	kill[attacked] = []EntityId{}
	return nil
}

// 助攻
func assistingKill(inner, args []interface{}) []interface{} {
	_, kill, _ := ParseAwardInner(inner)
	// 攻击方
	attack := EntityId(args[0].(string))
	// 被攻击方
	attacked := EntityId(args[1].(string))
	v, ok := kill[attacked]
	if !ok {
		kill[attacked] = []EntityId{attack}
		return nil
	}
	// 简单的存储规则：按时序递增(数组索引递增)保留最新的6条数据
	if len(v) >= 6 {
		for i, k := range v {
			if i == 0 {
				continue
			}
			v[i-1] = k
		}
		v[5] = attack
		return nil
	}
	v = append(v, attack)
	kill[attacked] = v
	return nil
}

// 格式化打印
func formatPrintAward(attack, attacked EntityId, list []string, a, b []int) string {
	return fmt.Sprintf("[%s:%s]被[%s:%s]干死了助攻方有%v,分别统计战绩[%v=>%v]", GetName(attacked), string(attacked),
		GetName(attack), string(attack), list, a, b)
}
