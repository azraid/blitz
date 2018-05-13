package main

import (
	"math/rand"
	"sync"
	"time"

	"github.com/azraid/pasque/app"

	. "github.com/azraid/blitz/services/juli"
	. "github.com/azraid/pasque/core"
)

type SingleInfo struct {
	objID   int
	dolKind TDol
	drawPos POS
}

func newSingleInfo() *SingleInfo {
	return &SingleInfo{}
}

type ServerBlock struct {
	SingleInfo
	grpID          int
	dolStat        TDStat
	posY           float64
	fallWaitTimeMs int64
}

func newServerBlock(objID int, pos POS) *ServerBlock {
	sb := &ServerBlock{
		SingleInfo: SingleInfo{
			objID:   objID,
			drawPos: pos,
			dolKind: EDOL_NORMAL_MAX,
		},
		grpID:          -1,
		dolStat:        EDSTAT_NONE,
		posY:           0,
		fallWaitTimeMs: 0,
	}

	return sb
}

type ServerGroup struct {
	grpID  int
	cnt    int
	blocks []*ServerBlock
}

func newServerGroup(grpID int) *ServerGroup {
	sg := &ServerGroup{
		grpID: grpID,
		cnt:   0,
	}

	sg.blocks = make([]*ServerBlock, 6)
	for k, _ := range sg.blocks {
		sg.blocks[k] = newServerBlock(-1, POS{X: -1, Y: -1})
	}

	return sg
}

type GameOption struct {
	responseDelayTimeMs int
	xsize               int
	xmax                int
	ysize               int
	ymax                int
	cnstOff             int
	cnstIdx             int
	cnsts               []TCnst
}

type GridData struct {
	p1       *Player
	p2       *Player
	opt      *GameOption
	GameStat TGStat
	Mode     TGMode
	lock     *sync.RWMutex
	tick     *time.Ticker
}

var procTimer time.Duration = time.Millisecond * DEFAULT_TICK_MS

func CreateGridData(key string, mode TGMode, gridData interface{}) (g *GridData) {
	if gridData != nil {
		g = gridData.(*GridData)
	} else {
		g = &GridData{GameStat: EGROOM_STAT_INIT, Mode: mode}
		g.lock = new(sync.RWMutex)
		g.tick = time.NewTicker(procTimer)
	}

	g.opt = &GameOption{
		responseDelayTimeMs: 0,
		xsize:               7,
		xmax:                6,
		ysize:               11,
		ymax:                10,
		cnstOff:             0,
		cnstIdx:             0,
	}

	g.opt.cnsts = append(g.opt.cnsts, ECNST_V3)
	g.opt.cnsts = append(g.opt.cnsts, ECNST_I2)
	g.opt.cnsts = append(g.opt.cnsts, ECNST_V3)
	g.opt.cnsts = append(g.opt.cnsts, ECNST_I2)
	g.opt.cnsts = append(g.opt.cnsts, ECNST_I3)
	g.opt.cnsts = append(g.opt.cnsts, ECNST_O4)
	g.opt.cnsts = append(g.opt.cnsts, ECNST_S4)
	g.opt.cnsts = append(g.opt.cnsts, ECNST_Z4)
	g.opt.cnsts = append(g.opt.cnsts, ECNST_J4)
	g.opt.cnsts = append(g.opt.cnsts, ECNST_L4)
	g.opt.cnsts = append(g.opt.cnsts, ECNST_V3)
	g.opt.cnsts = append(g.opt.cnsts, ECNST_I2)
	g.opt.cnsts = append(g.opt.cnsts, ECNST_I3)
	g.opt.cnsts = append(g.opt.cnsts, ECNST_V3)
	g.opt.cnsts = append(g.opt.cnsts, ECNST_I2)
	g.opt.cnsts = append(g.opt.cnsts, ECNST_I3)
	g.opt.cnsts = append(g.opt.cnsts, ECNST_O4)
	g.opt.cnsts = append(g.opt.cnsts, ECNST_S4)
	g.opt.cnsts = append(g.opt.cnsts, ECNST_Z4)
	g.opt.cnsts = append(g.opt.cnsts, ECNST_J4)
	g.opt.cnsts = append(g.opt.cnsts, ECNST_L4)
	g.opt.cnsts = append(g.opt.cnsts, ECNST_V3)
	g.opt.cnsts = append(g.opt.cnsts, ECNST_I2)
	g.opt.cnsts = append(g.opt.cnsts, ECNST_I3)

	// Shuffle

	for i, _ := range g.opt.cnsts {
		pick := rand.Intn(len(g.opt.cnsts) - 1)
		g.opt.cnsts[i], g.opt.cnsts[pick] = g.opt.cnsts[pick], g.opt.cnsts[i]
	}

	//////////////////////////////////////////////////////////////////////////
	// Padding Full Same Set
	g.opt.cnsts = append(g.opt.cnsts, g.opt.cnsts...)

	g.GameStat = EGROOM_STAT_INIT
	return g
}

func (g *GridData) SetPlayer(userID TUserID) (*Player, error) {
	if g.p1 == nil {
		g.p1 = newPlayer(userID, 1)
		return g.p1, nil
	}

	if g.p1.userID == userID {
		return g.p1, nil
	}

	if g.p2 == nil {
		g.p2 = newPlayer(userID, 2)
		g.p1.other = g.p2
		g.p2.other = g.p1
		return g.p2, nil
	}

	if g.p2.userID == userID {
		g.p1.other = g.p2
		g.p2.other = g.p1
		return g.p2, nil
	}

	return nil, IssueErrorf("UserID is not matched")
}

func (g *GridData) GetPlayer(userID TUserID) (*Player, error) {
	if g.p1 != nil && g.p1.userID == userID {
		return g.p1, nil
	}

	if g.p2 != nil && g.p2.userID == userID {
		return g.p2, nil
	}

	return nil, IssueErrorf("Not found Player")
}

func (g *GridData) PlayReady(userID TUserID) error {
	p, err := g.GetPlayer(userID)
	if err != nil {
		return err
	}
	p.Init(g.opt.xsize, g.opt.ysize)
	p.SetCnstList(g.opt.cnsts)

	return nil
}

func (g *GridData) RemovePlayer(userID TUserID) {
	if g.p1 != nil && g.p1.userID == userID {
		g.p1 = nil
		g.GameStat = EGROOM_STAT_END
		if g.p2 != nil {
			g.p2.stat = EPSTAT_INIT
		}

	} else if g.p2 != nil && g.p2.userID == userID {
		g.p2 = nil
		g.GameStat = EGROOM_STAT_END
		if g.p1 != nil {
			g.p1.stat = EPSTAT_INIT
		}
	}

	if g.p1 == nil && g.p2 == nil {
		g.Final()
	}
}

func (g *GridData) IsNull() bool {
	if g.p1 == nil && g.p2 == nil && g.tick == nil {
		return true
	}
	return false
}

func (g *GridData) TryStart() bool {
	if g.Mode == EGMODE_SP && g.p1.stat == EPSTAT_READY {
		go goPlay(g, time.Now())
		return true
	} else if g.Mode == EGMODE_PE && g.p1.stat == EPSTAT_READY {
		go goPlay(g, time.Now())
		return true
	} else if g.p1.stat == EPSTAT_READY && g.p2 != nil && g.p2.stat == EPSTAT_READY {
		go goPlay(g, time.Now())
		return true
	}
	return false
}

func (g *GridData) Final() {
	g.tick.Stop()
	g.tick = nil
}

func goPlay(g *GridData, beforeT time.Time) {
	defer app.DumpRecover()

	g.GameStat = EGROOM_STAT_READY

	if g.Mode == EGMODE_PP {
		SendPlayStart(g.p1.userID, g.p1)
		SendPlayStart(g.p2.userID, g.p2)
		g.p1.stat = EPSTAT_RUNNING
		g.p2.stat = EPSTAT_RUNNING
	} else {
		SendPlayStart(g.p1.userID, g.p1)
		g.p1.stat = EPSTAT_RUNNING
	}

	for _ = range g.tick.C {
		elapsed := time.Now().Sub(beforeT)
		if elapsed.Nanoseconds() > procTimer.Nanoseconds() {
			if gap := (elapsed.Nanoseconds() - procTimer.Nanoseconds()) / int64(time.Millisecond); gap > 100 {
				app.ErrorLog("-----------------Too Slow %d ms", gap)
			}
		}

		elapsedTimeMs := int(elapsed.Nanoseconds() / int64(time.Millisecond))
		g.Go(elapsedTimeMs)
		beforeT = time.Now()
	}
}

func (g *GridData) Go(elapsedTimeMs int) {
	g.Lock()

	defer func() {
		app.DumpRecover()
		g.Unlock()
	}()

	if g.GameStat != EGROOM_STAT_READY {
		return
	}

	g.GameStat = EGROOM_STAT_PLAYING

	var loser *Player
	if g.p1 != nil {
		if ok := g.p1.Play(int64(elapsedTimeMs), g.Mode); !ok {
			loser = g.p1
		}
	}

	if g.Mode == EGMODE_PP && g.p2 != nil {
		if ok := g.p2.Play(int64(elapsedTimeMs), g.Mode); !ok {
			loser = g.p2
		}
	}

	if loser != nil {
		SendCPlayEnd(loser.userID, loser, EEND_LKO)
		if loser.other != nil {
			SendCPlayEnd(loser.other.userID, loser, EEND_LKO)
		}

		g.GameStat = EGROOM_STAT_END
		g.Final()
		return
	}

	g.GameStat = EGROOM_STAT_READY

}

func (g *GridData) Lock() {
	g.lock.Lock()
}

func (g *GridData) Unlock() {
	g.lock.Unlock()
}
