package handlers

import (
	Msg "gserver_service_demo/src/server/protomsg"

	"github.com/gfandada/gserver/network"
	Service "github.com/gfandada/gserver/services/service"
)

func testReqHandler(args []interface{}) []interface{} {
	data := args[0].(*network.RawMessage).MsgData
	test := data.(*Msg.TestReq).Test
	msg := network.RawMessage{
		MsgId: uint16(2002),
		MsgData: &Msg.TestPush{
			Test: test,
		},
	}
	Service.ForEachSend(msg)
	return nil
}

func testPushHandler(args []interface{}) []interface{} {
	return nil
}
