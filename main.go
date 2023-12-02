package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/anoshenko/rui"

	"github.com/ruraomsk/ag-server/logger"
	"github.com/ruraomsk/ruod/controller"
	"github.com/ruraomsk/ruod/hardware"
	"github.com/ruraomsk/ruod/setup"
	"github.com/ruraomsk/ruod/web"
)

func init() {
	setup.Set = new(setup.Setup)
	setup.ExtSet = new(setup.ExtSetup)
	if _, err := toml.DecodeFS(resources, "config/base.toml", &setup.Set); err != nil {
		fmt.Println("Dismiss base.toml")
		os.Exit(-1)
		return
	}
	if _, err := os.Stat("config.json"); err == nil {
		file, err := os.ReadFile("config.json")
		if err == nil {
			err = json.Unmarshal(file, &setup.ExtSet)
			setup.Set.Update(*setup.ExtSet)
		}
	}
	setup.ExtSet.Update(*setup.Set)
	_ = os.MkdirAll(setup.Set.LogPath, 0777)
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if err := logger.Init(setup.Set.LogPath); err != nil {
		log.Panic("Error logger system", err.Error())
		return
	}
	fmt.Println("Ruod start")
	logger.Info.Println("Ruod start")
	go hardware.Start()
	time.Sleep(time.Second)
	go controller.Start()
	go rui.AddEmbedResources(&resources)
	go web.Web()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("\nwait ...")
	time.Sleep(1 * time.Second)
	fmt.Println("Ruod stop")
	logger.Info.Println("Ruod stop")
}
