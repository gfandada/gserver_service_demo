package match

import (
	"gserver_service_demo/src/server/protomsg"

	"github.com/gfandada/gserver/network"
	"github.com/gfandada/gserver/services/service"
	"github.com/gfandada/gserver/util"
	"github.com/golang/protobuf/proto"
)

func PostMatchSucceed(reds []int32, blus []int32) string {
	fightid := util.NewV4().String()
	for _, v := range reds {
		service.Send(v, network.RawMessage{
			MsgId: uint16(1501),
			MsgData: &protomsg.MatchPush{
				Fighttype:  proto.Int32(int32(len(reds))),
				Fightid:    proto.String(fightid),
				Playertype: proto.Int32(int32(protomsg.ENTITY_RED_SHIP_FLAG)),
				Reds:       reds,
				Blues:      blus,
			},
		})
	}
	for _, v := range blus {
		service.Send(v, network.RawMessage{
			MsgId: uint16(1501),
			MsgData: &protomsg.MatchPush{
				Fighttype:  proto.Int32(int32(len(blus))),
				Fightid:    proto.String(fightid),
				Playertype: proto.Int32(int32(protomsg.ENTITY_BLUE_SHIP_FLAG)),
				Reds:       reds,
				Blues:      blus,
			},
		})
	}
	return fightid
}
