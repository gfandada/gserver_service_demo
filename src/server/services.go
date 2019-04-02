package main

import (
	"sync"

	"github.com/gfandada/gserver"
	"github.com/gfandada/gserver/network"
	Msg "gserver_service_demo/src/server/protomsg"
	FightHandlers "gserver_service_demo/src/server/fight/handler"
	PortHandlers "gserver_service_demo/src/server/handlers"
	Match "gserver_service_demo/src/server/match"
	FightAi "gserver_service_demo/src/server/fight/entity"
)

var (
	once sync.Once
)

func RegisterServices() *network.MsgManager {
	once.Do(func() {
		//gserver.RunMysqlService("../../cfg/mysqllogger.xml", "../../cfg/db.json")
		gserver.RunXlsxService("C:/Users/Administrator/go/src/gserver_service_demo/cfg/xlsxlogger.xml", "C:/Users/Administrator/go/src/gserver_service_demo/cfg/xlsx/")
		PortHandlers.NewHandlers()
		FightHandlers.NewFightHandlers()
		Match.CreateMatch(int(Msg.TYPE_F1V1))
		Match.CreateMatch(int(Msg.TYPE_F3V3))
		Match.CreateMatch(int(Msg.TYPE_F5V5))
		FightAi.NewAiManager("../../cfg/aicfg")
		// Fight.NewFightManager()
	})
	return Msg.NewMsgCoder()
}
