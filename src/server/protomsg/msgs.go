package protomsg

import (
	"github.com/gfandada/gserver/network"
)

func NewMsgCoder() *network.MsgManager {
	coder := network.NewMsgManager()
	coder.Register(&network.RawMessage{
		MsgId:   uint16(2),
		MsgData: &ErrorAck{},
	})
	// for match
	coder.Register(&network.RawMessage{
		MsgId:   uint16(1500),
		MsgData: &MatchReq{},
	})
	coder.Register(&network.RawMessage{
		MsgId:   uint16(1501),
		MsgData: &MatchPush{},
	})
	// for test push
	coder.Register(&network.RawMessage{
		MsgId:   uint16(2000),
		MsgData: &TestReq{},
	})
	coder.Register(&network.RawMessage{
		MsgId:   uint16(2001),
		MsgData: &TestAck{},
	})
	coder.Register(&network.RawMessage{
		MsgId:   uint16(2002),
		MsgData: &TestPush{},
	})
	// for fight
	coder.Register(&network.RawMessage{
		MsgId:   uint16(2102),
		MsgData: &JoinFightReq{},
	})
	coder.Register(&network.RawMessage{
		MsgId:   uint16(2103),
		MsgData: &JoinFightAck{},
	})
	coder.Register(&network.RawMessage{
		MsgId:   uint16(2104),
		MsgData: &JoinFightPush{},
	})
	coder.Register(&network.RawMessage{
		MsgId:   uint16(2105),
		MsgData: &MoveReq{},
	})
	coder.Register(&network.RawMessage{
		MsgId:   uint16(2106),
		MsgData: &MovePush{},
	})
	coder.Register(&network.RawMessage{
		MsgId:   uint16(2107),
		MsgData: &AttackPush{},
	})
	coder.Register(&network.RawMessage{
		MsgId:   uint16(2108),
		MsgData: &SpwanPush{},
	})
	coder.Register(&network.RawMessage{
		MsgId:   uint16(2109),
		MsgData: &DeadPush{},
	})
	coder.Register(&network.RawMessage{
		MsgId:   uint16(2110),
		MsgData: &OverPush{},
	})
	return coder
}
