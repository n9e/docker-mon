package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/n9e/docker-mon/collect"
	"github.com/n9e/docker-mon/config"

	"github.com/toolkits/pkg/logger"
)

func InitLogger(l config.LoggerSection) {
	lb, err := logger.NewFileBackend(l.Dir)
	if err != nil {
		fmt.Println("cannot init logger:", err)
		os.Exit(1)
	}

	lb.SetRotateByHour(true)
	lb.SetKeepHours(l.KeepHours)

	logger.SetLogging(l.Level, lb)
}

func init() {
	cfg := flag.String("c", "docker-mon.yml", "configuration file")
	version := flag.Bool("v", false, "show version")
	flag.Parse()

	handleVersion(*version)
	handleConfig(*cfg)
	InitLogger(config.G.Logger)
}

func main() {
	collect.HostIp = collect.HostIpMap{M: make(map[string]string)}
	hostIpFile := ".host_ip"
	collect.HostIp.Load(hostIpFile)

	collect.Collect()
	err := collect.HostIp.WriteFile(hostIpFile)
	if err != nil {
		logger.Error(err)
	}

	logger.Close()
}

func handleVersion(displayVersion bool) {
	if displayVersion {
		fmt.Println(config.Version)
		os.Exit(0)
	}
}

func handleConfig(configFile string) {
	err := config.Parse(configFile)
	if err != nil {
		log.Fatalln(err)
	}
}
