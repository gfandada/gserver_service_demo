package fightNew

import (
	. "github.com/magicsea/behavior3go/core"
)

type Ai struct {
	Tree  *BehaviorTree
	Black *Blackboard
}

// 属性常量定义
const (
	HP      = "hp"      // 当前耐久
	HP_OLD  = "hp_old"  // 原先耐久
	LEVEL   = "level"   // 当前等级
	EXP     = "exp"     // 当前经验
	DEFENSE = "defense" // 当前防御
)

// fightScheduler
const (
	FLUSH_SOLDIER = "flush_soldier" // 刷新小兵
	JOIN_FIGHT    = "join_fight"    // 加入战斗
	GET_FIGHT     = "get_fight"     // 获取战斗双方数据
	LEAVE_FIGHT   = "leave_fight"   // 离开战斗
	MOVE_FIGHT    = "move_fight"    // 移动指令
	GET_AOIS      = "get_aois"      // 获取aoi信息
	ATTACK_ADAP   = "attack_adap"   // 攻击适配
	ATTACK_KILL   = "attack_kill"   // 击杀
	ADD_EQUIP     = "add_equip"     // 添加装备
	RM_EQUIP      = "rm_equip"      // 移除装备
	ROADCAST_TMP  = "broadcast_tmp" // 通用的战斗场景内广播
)

// fightdamage
const (
	AUTO_ATTACK = "auto_attack" // 自动攻击
)

// fightaward
const (
	KILL       = "kill"          // 击杀
	ASSISTKILL = "assistingkill" // 助攻
)

// fightpost
const (
	ENTER_POST        = "enter_post"        // 玩家加入推送
	AOI_POST          = "aoi_post"          // aoi推送
	SPWAN_ENTITY_POST = "spwan_entity_post" // 实体派生推送
	DEATH_ENTITY_POST = "death_entity_post" // 实体死亡推送
	GAME_OVER_POST    = "game_over_post"    // 游戏结束推送
	BROADCAST         = "broadcast"         // 通用的战斗场景内广播
)
