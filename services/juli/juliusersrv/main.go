package main

import (
	"fmt"
	"os"

	. "github.com/azraid/blitz/services/juli"
	"github.com/azraid/pasque/app"
	n "github.com/azraid/pasque/core/net"
)

const GameSpn = "Julivonoblitz.Tcgate"

func main() {

	if len(os.Args) < 2 {
		fmt.Println("ex) juliusersrv.exe [eid]")
		os.Exit(1)
	}

	eid := os.Args[1]

	workPath := "./"
	if len(os.Args) == 3 {
		workPath = os.Args[2]
	}

	app.InitApp(eid, "", workPath)

	cli := n.NewClient(eid)
	cli.RegisterGridHandler(n.GetNameOfApiMsg(JoinInMsg{}), OnJoinIn)
	cli.RegisterGridHandler(n.GetNameOfApiMsg(PlayReadyMsg{}), OnPlayReady)
	cli.RegisterGridHandler(n.GetNameOfApiMsg(CPlayStartMsg{}), OnCPlayStart)
	cli.RegisterGridHandler(n.GetNameOfApiMsg(CPlayEndMsg{}), OnCPlayEnd)
	cli.RegisterGridHandler(n.GetNameOfApiMsg(CMatchUpMsg{}), OnCMatchUp)
	cli.RegisterGridHandler(n.GetNameOfApiMsg(LeaveRoomMsg{}), OnLeaveRoom)
	cli.RegisterGridHandler(n.GetNameOfApiMsg(DrawGroupMsg{}), OnDrawGroup)
	cli.RegisterGridHandler(n.GetNameOfApiMsg(DrawSingleMsg{}), OnDrawSingle)

	toplgy := n.Topology{
		Spn:           app.Config.Spn,
		FederatedKey:  "UserID",
		FederatedApis: cli.ListGridApis()}

	cli.Dial(toplgy)

	app.WaitForShutdown()
	return
}
