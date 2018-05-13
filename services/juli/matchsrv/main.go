package main

import (
	"fmt"
	"os"

	. "github.com/azraid/blitz/services/juli"
	"github.com/azraid/pasque/app"
	n "github.com/azraid/pasque/core/net"
)

var rpcx n.Client

func main() {

	if len(os.Args) < 2 {
		fmt.Println("ex) matchsrv.exe [eid]")
		os.Exit(1)
	}

	eid := os.Args[1]

	workPath := "./"
	if len(os.Args) == 3 {
		workPath = os.Args[2]
	}

	app.InitApp(eid, "", workPath)

	rpcx = n.NewClient(eid)
	rpcx.RegisterRandHandler(n.GetNameOfApiMsg(MatchPlayMsg{}), OnMatchPlay)
	rpcx.RegisterRandHandler(n.GetNameOfApiMsg(LeaveWaitingMsg{}), OnLeaveWaiting)

	toplgy := n.Topology{Spn: app.Config.Spn}

	rpcx.Dial(toplgy)

	app.WaitForShutdown()
	return
}
