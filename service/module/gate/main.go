package main

import (
	"flag"
	"fmt"
	// "os"
	"path/filepath"
	"zest/engine/common"
	"zest/engine/conf"
	"zest/engine/zslog"
	"zest/service/sys"
)

func parseArgs() {
	var (
		sidint     int
		configFile string
		service    string
	)
	flag.IntVar(&sidint, "sid", 0, "set sid")
	flag.StringVar(&configFile, "configfile", "", "set service config file")
	flag.StringVar(&service, "service", "default", "set service module")
	flag.Parse()
	sid := int32(sidint)
	conf.SetGlobalConf("Sid", sid)
	conf.SetGlobalConf("configfile", configFile)
	conf.SetGlobalConf("Service", service)
}

// 预处理函数,加载配置,设置日志,启动服务
func perProcess() {
	parseArgs()
	fmt.Printf("module %v sid %v\n", conf.GetGlobalConf("Service").(string), conf.GetGlobalConf("Sid").(int32))
	conf.SetConfigFile(conf.GetGlobalConf("configfile").(string))
	conf.LoadConfig()
	fmt.Println("confilg load success")
	Path, err := common.GetProgrammePath()
	if err != nil {
		fmt.Println(err)
	}
	LogDir := filepath.Join(filepath.Dir(filepath.Dir(Path)), "log")
	zslog.SetupLog(LogDir, fmt.Sprintf("%v_%v", conf.GetGlobalConf("Service").(string), conf.GetGlobalConf("Sid").(int32)), zslog.ParseLevel("debug"))
	fmt.Println("log system load success")
	zslog.LogDebug("log system load success")

	s := sys.NewService(conf.GetGlobalConf("Sid").(int32), conf.GetGlobalConf("Service").(string))
	s.Start()
	fmt.Println(s)
}

func main() {
	perProcess()
	NewGate(sys.GetService().GetOutServer().GetConnMgr())
	select {}
}
