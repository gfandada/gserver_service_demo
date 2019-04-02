package match

import (
	"time"

	. "gserver_service_demo/src/server/fight/handler"

	. "github.com/gfandada/gserver/gameutil/fight"
	"github.com/gfandada/gserver/logger"
)

type MatchScheduler struct {
	Flag        int
	waitingblue []int32
	waitingred  []int32
}

func (m *MatchScheduler) Name() string {
	return GetMatchName(m.Flag)
}

func (m *MatchScheduler) SetTimer() time.Duration {
	return 0
}

func (m *MatchScheduler) TimerWork() {
}

func (m *MatchScheduler) Init() {
}

func (m *MatchScheduler) Close() {
}

func (m *MatchScheduler) Handler(msg string, args []interface{}, ret chan []interface{}) {
	switch msg {
	case MATCHING:
		m.matching(args)
	}
}

func (m *MatchScheduler) matching(args []interface{}) {
	userid := args[0].(int32)
	if len(m.waitingblue) < m.Flag {
		m.waitingblue = append(m.waitingblue, userid)
		return
	}
	if len(m.waitingred) < m.Flag {
		m.waitingred = append(m.waitingred, userid)
		if len(m.waitingred) >= m.Flag {
			logger.Debug("%s over blues %v reds %v", m.Name(), m.waitingblue, m.waitingred)
			fightid := PostMatchSucceed(m.waitingblue, m.waitingred)
			if err := CreateFight(FightId(fightid), m.Flag, new(FightMap)); err != nil {
				logger.Error("%s over CreateFight %s error %v", m.Name(), fightid, err)
			}
			logger.Info("CreateFight %s fightid %s blues %v reds %v", m.Name(), fightid,
				m.waitingblue, m.waitingred)
			m.waitingblue = []int32{}
			m.waitingred = []int32{}
		}
	}
}
