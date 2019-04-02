package handler

import (
	"strconv"

	. "gserver_service_demo/src/server/fight/entity"
	. "gserver_service_demo/src/server/protomsg"

	. "github.com/gfandada/gserver/gameutil/entity"
	"github.com/gfandada/gserver/network"
	"github.com/golang/protobuf/proto"
)

// 玩家加入推送
func NewEnterPush(entity *Entity) network.RawMessage {
	typ := entity.I.(*Ship).Flag()
	entityid := string(entity.Id)
	ms := NewAoi(entity.GetPosition())
	userid := entity.Client.GetId()
	return network.RawMessage{
		MsgId: 2104,
		MsgData: &JoinFightPush{
			Spawn: &SpawnState{
				Type:     proto.Int32(typ),
				Entityid: proto.String(entityid),
				Ms:       ms,
				UserId:   proto.Int32(userid),
			},
		},
	}
}

// move-push
func NewAoiPush(ships map[*Entity]struct{}, soldiers map[*Entity]struct{}) network.RawMessage {
	var data []*AOIPushItem
	for entity := range ships {
		entityid := string(entity.Id)
		ms := NewAoi(entity.GetPosition())
		data = append(data, &AOIPushItem{
			Entityid: proto.String(entityid),
			Ms:       ms,
		})
	}
	for entity := range soldiers {
		entityid := string(entity.Id)
		ms := NewAoi(entity.GetPosition())
		data = append(data, &AOIPushItem{
			Entityid: proto.String(entityid),
			Ms:       ms,
		})
	}
	return network.RawMessage{
		MsgId: 2106,
		MsgData: &MovePush{
			Aoi: data,
		},
	}
}

func NewAttackPush(userid1, userid2 string, attack, hp int32) network.RawMessage {
	return network.RawMessage{
		MsgId: 2107,
		MsgData: &AttackPush{
			Attack:   proto.String(userid1),
			Attacked: proto.String(userid2),
			Damage:   proto.Int32(attack),
			Hp:       proto.Int32(hp),
		},
	}
}

func NewAoi(aoi Vector3) *AOI {
	return &AOI{
		Time: proto.Int64(aoi.TIME),
		Pos: &Vector2{
			X: proto.Float32(float32(aoi.X)),
			Y: proto.Float32(float32(aoi.Z)),
		},
		V: &Vector2{
			X: proto.Float32(float32(aoi.VX)),
			Y: proto.Float32(float32(aoi.VZ)),
		},
	}
}

// 新建派生的实体数据
func NewSpawnState(entitys map[*Entity]struct{}) (ret []*SpawnState) {
	for entity := range entitys {
		res, _ := strconv.Atoi(entity.Desc.Name)
		switch entity.Desc.Flag {
		case int32(ENTITY_BLUE_SHIP_FLAG):
			ret = append(ret, &SpawnState{
				ResId:    proto.Int32(int32(res)),
				Type:     proto.Int32(entity.I.(*Ship).Flag()),
				Entityid: proto.String(string(entity.Id)),
				Ms:       NewAoi(entity.GetPosition()),
				UserId:   proto.Int32(entity.Client.GetId()),
			})
		case int32(ENTITY_RED_SHIP_FLAG):
			ret = append(ret, &SpawnState{
				ResId:    proto.Int32(int32(res)),
				Type:     proto.Int32(entity.I.(*Ship).Flag()),
				Entityid: proto.String(string(entity.Id)),
				Ms:       NewAoi(entity.GetPosition()),
				UserId:   proto.Int32(entity.Client.GetId()),
			})
		case int32(ENTITY_RED_SOLDIER_FLAG):
			ret = append(ret, &SpawnState{
				ResId:    proto.Int32(int32(res)),
				Type:     proto.Int32(entity.I.(*Soldier).Flag()),
				Entityid: proto.String(string(entity.Id)),
				Ms:       NewAoi(entity.GetPosition()),
			})
		case int32(ENTITY_BLUE_SOLDIER_FLAG):
			ret = append(ret, &SpawnState{
				ResId:    proto.Int32(int32(res)),
				Type:     proto.Int32(entity.I.(*Soldier).Flag()),
				Entityid: proto.String(string(entity.Id)),
				Ms:       NewAoi(entity.GetPosition()),
			})
		case int32(ENTITY_BLUE_TOWER_FLAG):
			ret = append(ret, &SpawnState{
				ResId:    proto.Int32(int32(res)),
				Type:     proto.Int32(entity.I.(*Tower).Flag()),
				Entityid: proto.String(string(entity.Id)),
				Ms:       NewAoi(entity.GetPosition()),
			})
		case int32(ENTITY_RED_TOWER_FLAG):
			ret = append(ret, &SpawnState{
				ResId:    proto.Int32(int32(res)),
				Type:     proto.Int32(entity.I.(*Tower).Flag()),
				Entityid: proto.String(string(entity.Id)),
				Ms:       NewAoi(entity.GetPosition()),
			})
		case int32(ENTITY_RED_CRYSTAL):
			ret = append(ret, &SpawnState{
				ResId:    proto.Int32(int32(res)),
				Type:     proto.Int32(entity.I.(*Crystal).Flag()),
				Entityid: proto.String(string(entity.Id)),
				Ms:       NewAoi(entity.GetPosition()),
			})
		case int32(ENTITY_BLUE_CRYSTAL):
			ret = append(ret, &SpawnState{
				ResId:    proto.Int32(int32(res)),
				Type:     proto.Int32(entity.I.(*Crystal).Flag()),
				Entityid: proto.String(string(entity.Id)),
				Ms:       NewAoi(entity.GetPosition()),
			})
		}
	}
	return
}

// 实体出生
func NewSpwanPush(entitys map[*Entity]struct{}) network.RawMessage {
	return network.RawMessage{
		MsgId: 2108,
		MsgData: &SpwanPush{
			Spwans: NewSpawnState(entitys),
		},
	}
}

// 实体死亡
func NewSpwanDeadPush(dead, kill string, assists []string, damage int32) network.RawMessage {
	return network.RawMessage{
		MsgId: 2109,
		MsgData: &DeadPush{
			Death: []*DeadSettlement{
				&DeadSettlement{
					Dead:    proto.String(dead),
					Kill:    proto.String(kill),
					Assists: assists,
					Damage:  proto.Int32(damage),
				},
			},
		},
	}
}

// 游戏结束
func NewGameOverPush(flag int32, statistics map[EntityId][]int) network.RawMessage {
	var reds, blues []*EndSettlement
	for k := range statistics {
		if entity := GetEntity(k); entity != nil {
			if entity.Desc.Flag == int32(ENTITY_BLUE_SHIP_FLAG) {
				blues = append(blues, &EndSettlement{
					Entityid: proto.String(string(k)),
					Kill:     proto.Int32(int32(statistics[k][0])),
					Dead:     proto.Int32(int32(statistics[k][1])),
					Assists:  proto.Int32(int32(statistics[k][2])),
				})
			} else if entity.Desc.Flag == int32(ENTITY_RED_SHIP_FLAG) {
				reds = append(reds, &EndSettlement{
					Entityid: proto.String(string(k)),
					Kill:     proto.Int32(int32(statistics[k][0])),
					Dead:     proto.Int32(int32(statistics[k][1])),
					Assists:  proto.Int32(int32(statistics[k][2])),
				})
			}
		}
	}
	return network.RawMessage{
		MsgId: 2110,
		MsgData: &OverPush{
			Camp: proto.Int32(flag),
			Red:  reds,
			Blue: blues,
		},
	}
}
