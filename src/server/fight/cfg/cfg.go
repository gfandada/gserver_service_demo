package cfg

import (
	"strconv"

	. "github.com/gfandada/gserver/loader"
)

// 获取战船耐久、防御
func GetShipDura(level uint32) (durability, defend uint32) {
	loader := new(Loader)
	// 耐久
	durability, _ = loader.GetUint32(loader.Get("ShipLevel", level, "Durability"))
	// 防御
	defend, _ = loader.GetUint32(loader.Get("ShipLevel", level, "DefendPoint"))
	return
}

// 获取战船攻击、攻击范围、攻击间隔
func GetShipAttack(level uint32) (attack, radius, duration uint32) {
	loader := new(Loader)
	// 攻击力
	attack, _ = loader.GetUint32(loader.Get("ShipLevel", level, "AttackPoint"))
	// 攻击范围
	radius, _ = loader.GetUint32(loader.Get("ShipLevel", level, "AttackRadius"))
	// 攻击间隔
	duration, _ = loader.GetUint32(loader.Get("ShipLevel", level, "AttackDuration"))
	return
}

// 获取防御塔(含水晶)耐久、防御
// 1-16：tower 17-18：crystal
func GetTowerDura(id uint32) (durability, defend uint32) {
	loader := new(Loader)
	// 耐久
	durability, _ = loader.GetUint32(loader.Get("Building", id, "Durability"))
	// 防御
	defend, _ = loader.GetUint32(loader.Get("Building", id, "DefendPoint"))
	return
}

// 获取防御塔(含水晶)攻击、攻击范围、攻击间隔
// 1-16：tower 17-18：crystal
func GetTowerAttack(id uint32) (attack, radius, duration uint32) {
	loader := new(Loader)
	// 攻击力
	attack, _ = loader.GetUint32(loader.Get("Building", id, "AttackPoint"))
	// 攻击范围
	radius, _ = loader.GetUint32(loader.Get("Building", id, "AttackRadius"))
	// 攻击间隔
	duration, _ = loader.GetUint32(loader.Get("Building", id, "AttackDuration"))
	return
}

// 获取小兵耐久、防御
func GetSoldierDura(id uint32) (durability, defend uint32) {
	loader := new(Loader)
	// 耐久
	durability, _ = loader.GetUint32(loader.Get("NPC", id, "Durability"))
	// 防御
	defend, _ = loader.GetUint32(loader.Get("NPC", id, "DefendPoint"))
	return
}

// 获取小兵攻击、攻击范围、攻击间隔
func GetSoldierAttack(id uint32) (attack, radius, duration uint32) {
	loader := new(Loader)
	// 攻击力
	attack, _ = loader.GetUint32(loader.Get("NPC", id, "AttackPoint"))
	// 攻击范围
	radius, _ = loader.GetUint32(loader.Get("NPC", id, "AttackRadius"))
	// 攻击间隔
	duration, _ = loader.GetUint32(loader.Get("NPC", id, "AttackDuration"))
	return
}

// 获取船只默认耐久、防御
func GetDefaultShip(id uint32) (durability, defend uint32) {
	loader := new(Loader)
	// 耐久
	durability, _ = loader.GetUint32(loader.Get("Ship1", id, "Durability"))
	// 防御
	defend, _ = loader.GetUint32(loader.Get("Ship1", id, "DefendPoint"))
	return
}

// 获取船只默认装备
func GetDefaultShipEquip(id uint32) []uint32 {
	loader := new(Loader)
	var equips []uint32
	for i := 1; i <= 6; i++ {
		id, _ := loader.GetUint32(loader.Get("Ship", id, "Equipment"+strconv.Itoa(i)))
		equips = append(equips, id)
	}
	return equips
}

// 获取装备耐久、防御
func GetEquipDura(id uint32) (durability, defend uint32) {
	loader := new(Loader)
	// 耐久
	durability, _ = loader.GetUint32(loader.Get("Equipment", id, "Durability"))
	// 防御
	defend, _ = loader.GetUint32(loader.Get("Equipment", id, "DefendPoint"))
	return
}

// 获取装备攻击、攻击范围、攻击间隔
func GetEquipAttack(id uint32) (attack, radius, duration uint32) {
	loader := new(Loader)
	// 攻击力
	attack, _ = loader.GetUint32(loader.Get("Equipment", id, "Attack"))
	// 攻击范围
	radius, _ = loader.GetUint32(loader.Get("Equipment", id, "AttackRadius"))
	// 攻击间隔
	duration, _ = loader.GetUint32(loader.Get("Equipment", id, "AttackDuration"))
	return
}
