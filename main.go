package main

import (
	"time"

	"github.com/Impisigmatus/service_core/log"
)

func init() {
	log.Init(log.LevelDebug, "logs/debug.json")
}

func main() {
	log.Debugf(`{ "date": "%s" }`, time.Now().Format(time.DateTime))
}
