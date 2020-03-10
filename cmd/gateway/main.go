package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
	"github.com/devopsfaith/krakend/proxy"
	"github.com/devopsfaith/krakend/router/gin"
)

func main() {
	port := flag.Int("p", 8080, "Port of the service")
	logLevel := flag.String("l", "DEBUG", "Logging level")
	debug := flag.Bool("d", true, "Enable the debug")
	configFile := flag.String("c", "config.json", "Path to the configuration filename")
	flag.Parse()

	parser := config.NewParser()
	serviceConfig, err := parser.Parse(*configFile)
	if err != nil {
		fmt.Print(err.Error())
	}
	serviceConfig.Debug = serviceConfig.Debug || *debug
	if *port != 0 {
		serviceConfig.Port = *port
	}

	logger, _ := logging.NewLogger(*logLevel, os.Stdout, "[KRAKEND]")

	routerFactory := gin.DefaultFactory(proxy.DefaultFactory(logger), logger)

	routerFactory.New().Run(serviceConfig)
}
