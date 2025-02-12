package main

import (
	"github.com/ddev/ddev/cmd/ddev/cmd"
	"github.com/ddev/ddev/pkg/util"
	"os"
	"time"
)

func main() {
	runTime := util.TimeTrack(time.Now(), "main()")
	defer runTime()

	// Prevent running as root for most cases
	// We really don't want ~/.ddev to have root ownership, breaks things.
	if os.Geteuid() == 0 && len(os.Args) > 1 && os.Args[1] != "hostname" {
		util.Failed("ddev is not designed to be run with root privileges, please run as normal user and without sudo")
	}

	cmd.Execute()
}
