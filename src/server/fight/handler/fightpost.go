package handler

import (
	. "github.com/gfandada/gserver/gameutil/entity"
	"github.com/gfandada/gserver/network"
	"github.com/gfandada/gserver/services/service"
)

// 将进入的玩家信息推送给others
// @params entity:进入的玩家 others:待推送方
func PostEnter(entity *Entity, others map[*Entity]struct{}) {
	if len(others) <= 0 {
		return
	}
	for entity1 := range others {
		service.Send(entity1.Client.GetId(), NewEnterPush(entity))
	}
}

// 将others的aoi位置推送给userid
// @params time:时间戳ms ships:玩家 soldiers:小兵
func PostAoi(ships map[*Entity]struct{}, soldiers map[*Entity]struct{}) {
	if len(ships) <= 0 {
		return
	}
	for ship := range ships {
		if ship == nil {
			continue
		}
		service.Send(ship.Client.GetId(), NewAoiPush(ships, soldiers))
	}
}

// 推送实体派生
func PostEntitySpwan(ships map[*Entity]struct{}, others map[*Entity]struct{}) {
	for ship := range ships {
		service.Send(ship.Client.GetId(), NewSpwanPush(others))
	}
}

// 推送广播
func PostBroadcast(ships map[*Entity]struct{}, data interface{}) {
	for ship := range ships {
		service.Send(ship.Client.GetId(), data.(network.RawMessage))
	}
}

// 实体死亡推送
func PostEntityDead(ships map[*Entity]struct{}, dead, kill string, assists []string, damage int32) {
	for ship := range ships {
		service.Send(ship.Client.GetId(), NewSpwanDeadPush(dead, kill, assists, damage))
	}
}
