package handlers

import (
	Services "github.com/gfandada/gserver/services"
)

func NewHandlers() {
	Services.Register(uint16(Services.CLOSE_CONNECT), closeHandler)
	Services.Register(uint16(1500), matchReqHandler)
	Services.Register(uint16(2000), testReqHandler)
	Services.Register(uint16(2002), testPushHandler)
	Services.Register(uint16(2102), joinFightHandler)
	Services.Register(uint16(2105), moveFightHandler)
}
