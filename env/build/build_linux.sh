#!/bin/bash

if [ ! -d "$GOPATH/bin/blitz" ]; then
   mkdir $GOPATH/bin/blitz
fi

if [ ! -d "$GOPATH/bin/blitz/linux_amd64" ]; then
   mkdir $GOPATH/bin/blitz/linux_amd64
fi

if [ ! -d "$GOPATH/bin/blitz/linux_amd64/config" ]; then
   mkdir $GOPATH/bin/blitz/linux_amd64/config
fi

go build --o $GOPATH/bin/blitz/linux_amd64/spawn $GOPATH/src/github.com/azraid/pasque/bus/spawn/main.go 
go build --o $GOPATH/bin/blitz/linux_amd64/logsrv $GOPATH/src/github.com/azraid/pasque/bus/logsrv/main.go 

go build  -race --o $GOPATH/bin/blitz/linux_amd64/router $GOPATH/src/github.com/azraid/pasque/bus/router/main.go $GOPATH/src/github.com/azraid/pasque/bus/router/router.go
go build  -race --o $GOPATH/bin/blitz/linux_amd64/sgate $GOPATH/src/github.com/azraid/pasque/bus/sgate/main.go $GOPATH/src/github.com/azraid/pasque/bus/sgate/gate.go
go build -race -o $GOPATH/bin/blitz/linux_amd64/tcgate $GOPATH/src/github.com/azraid/pasque/bus/tcgate/main.go $GOPATH/src/github.com/azraid/pasque/bus/tcgate/gate.go $GOPATH/src/github.com/azraid/pasque/bus/tcgate/stub.go

go build  -race --o $GOPATH/bin/blitz/linux_amd64/sesssrv $GOPATH/src/github.com/azraid/pasque/services/auth/sesssrv/main.go $GOPATH/src/github.com/azraid/pasque/services/auth/sesssrv/db.go  $GOPATH/src/github.com/azraid/pasque/services/auth/sesssrv/grid.go  $GOPATH/src/github.com/azraid/pasque/services/auth/sesssrv/txn.go 
go build -race -o $GOPATH/bin/blitz/linux_amd64/juliworldsrv $GOPATH/src/github.com/azraid/blitz/services/juli/juliworldsrv/main.go $GOPATH/src/github.com/azraid/blitz/services/juli/juliworldsrv/grid.go  $GOPATH/src/github.com/azraid/blitz/services/juli/juliworldsrv/intxn.go  $GOPATH/src/github.com/azraid/blitz/services/juli/juliworldsrv/outtxn.go $GOPATH/src/github.com/azraid/blitz/services/juli/juliworldsrv/player.go
go build  -race --o $GOPATH/bin/blitz/linux_amd64/juliusersrv $GOPATH/src/github.com/azraid/blitz/services/juli/juliusersrv/main.go $GOPATH/src/github.com/azraid/blitz/services/juli/juliusersrv/grid.go  $GOPATH/src/github.com/azraid/blitz/services/juli/juliusersrv/txn.go
go build  -race --o $GOPATH/bin/blitz/linux_amd64/matchsrv $GOPATH/src/github.com/azraid/blitz/services/juli/matchsrv/main.go $GOPATH/src/github.com/azraid/blitz/services/juli/matchsrv/match.go $GOPATH/src/github.com/azraid/blitz/services/juli/matchsrv/txn.go 
go build -o $GOPATH/bin/blitz/linux_amd64/juli $GOPATH/src/github.com/azraid/blitz/test/juli/main.go  $GOPATH/src/github.com/azraid/blitz/test/juli/conn.go $GOPATH/src/github.com/azraid/blitz/test/juli/dialer.go $GOPATH/src/github.com/azraid/blitz/test/juli/resq.go $GOPATH/src/github.com/azraid/blitz/test/juli/client.go $GOPATH/src/github.com/azraid/blitz/test/juli/biz_login.go  $GOPATH/src/github.com/azraid/blitz/test/juli/biz_juli.go  $GOPATH/src/github.com/azraid/blitz/test/juli/cmd.go 

cp -rf $GOPATH/src/github.com/azraid/blitz/env/config/system_linux.json $GOPATH/bin/blitz/linux_amd64/config/system.json
cp -rf $GOPATH/src/github.com/azraid/blitz/env/run/run_linux.sh $GOPATH/bin/blitz/linux_amd64/run.sh
cp -rf $GOPATH/src/github.com/azraid/blitz/env/run/sampling.sh $GOPATH/bin/blitz/linux_amd64/sampling.sh
cp -rf $GOPATH/src/github.com/azraid/blitz/env/run/sampling10.sh $GOPATH/bin/blitz/linux_amd64/sampling10.sh
cp -rf $GOPATH/src/github.com/azraid/blitz/env/config/userauthdb.json $GOPATH/bin/blitz/linux_amd64/config/userauthdb.json
