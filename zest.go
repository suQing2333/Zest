package main

import (
	"flag"
	"fmt"
	// "os"
	"zest/engine/common"
	"zest/engine/conf"
	// "zest/engine/zslog"
)

var args struct {
	configFile string
	module     string
	sid        uint16
}

func parseArgs() {
	var sid int
	flag.IntVar(&sid, "sid", 0, "set sid")
	flag.StringVar(&args.configFile, "configfile", "", "set service config file")
	flag.StringVar(&args.module, "module", "", "set service module")
	flag.Parse()
	args.sid = uint16(sid)
}

func main() {
	fmt.Println("main run")
	parseArgs()
	conf.SetConfigFile(args.configFile)
	conf.LoadConfig()
	Path, err := common.GetProgrammePath()

}
