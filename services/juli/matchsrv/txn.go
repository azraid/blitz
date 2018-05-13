package main

import (
	"encoding/json"

	"github.com/azraid/pasque/app"
	//. "github.com/azraid/pasque/core"
	. "github.com/azraid/blitz/services/juli"
	n "github.com/azraid/pasque/core/net"
)

func OnMatchPlay(cli n.Client, req *n.RequestMsg) {
	var body MatchPlayMsg
	if err := json.Unmarshal(req.Body, &body); err != nil {
		app.ErrorLog(err.Error())
		cli.SendResWithError(req, RaiseNError(n.NErrorParsingError), nil)
		return
	}

	partner, ok := MatchPlayer(&Player{userID: body.UserID, grade: body.Grade})
	if ok {
		cli.SendRes(req, MatchPlayMsgR{OwnerID: partner, GuestID: body.UserID})
	} else {
		cli.SendRes(req, MatchPlayMsgR{})
	}
}

func OnLeaveWaiting(cli n.Client, req *n.RequestMsg) {
	var body MatchPlayMsg
	if err := json.Unmarshal(req.Body, &body); err != nil {
		app.ErrorLog(err.Error())
		cli.SendResWithError(req, RaiseNError(n.NErrorParsingError), nil)
		return
	}

	DeletePlayer(body.UserID)
	cli.SendRes(req, LeaveWaitingMsgR{})
}
