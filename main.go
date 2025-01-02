package main

import (
	_ "originserver2/authservice"
	_ "originserver2/centerservice"
	_ "originserver2/dbservice"
	_ "originserver2/gameservice"
	_ "originserver2/gateservice"
	_ "originserver2/loginservice"
	_ "originserver2/testservice"
	"time"

	"github.com/duanhf2012/origin/v2/node"
)

func main() {
	node.OpenProfilerReport(time.Second * 10)
	node.Start()
}
