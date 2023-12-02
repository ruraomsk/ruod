package controller

import (
	"time"

	"github.com/ruraomsk/ag-server/logger"
	"github.com/ruraomsk/ruod/controller/transport"
	"github.com/ruraomsk/ruod/hardware"
)

var live chan any
var rphase RepPhase
var rplan RepPlan
var rmajor RepMajor
var ralarm RepAlarm
var rsg RepSignalGroups
var rstatus RepStatus
var rsource RepSource

func getDuration() time.Duration {
	return 60 * time.Second
}

func controlCentral() {
	var timer *time.Timer
	hardware.SetWorkCentral <- 0
	logger.Info.Print("Нет управления от центра")
	for {
		<-live
		hardware.SetWorkCentral <- 1
		timer = time.NewTimer(getDuration())
		logger.Info.Print("Есть управление от центра")
		hardware.SetWorkCentral <- 1
	loop:
		for {
			select {
			case <-timer.C:
				hardware.SetWorkCentral <- 0
				break loop
			case <-live:
				timer.Stop()
				timer = time.NewTimer(getDuration())
			}
		}
		logger.Error.Print("Потеряно управление от центра")
	}

}

//Команды для КДМ
// case 0:	//Отключить управление
// case 1: //Переход в локальный режим
// case 2: //Переход в  режим ЖМ
// case 3: //Переход в  режим КК
// case 4: //Переход в  режим ОС
// case 5: //Хочет включить план координации
// case 6: //Хочет включить фазу

func Start() {
	go transport.Transport()
	live = make(chan any)
	go controlCentral()
	for {
		select {
		case command := <-transport.Commander:
			live <- 0
			logger.Debug.Printf("command %v", command)
			switch command.Code {
			case transport.CodeCallAllRed:
				hardware.CommandCentral(3, command.Value)
			case transport.CodeCallFlash:
				hardware.CommandCentral(2, command.Value)
			case transport.CodeCallDark:
				hardware.CommandCentral(4, command.Value)
			case transport.CodeCallPlan:
				hardware.CommandCentral(5, command.Value)
			case transport.CodeCallPhase:
				hardware.CommandCentral(6, command.Value)
			default:
				logger.Error.Printf("not command %v", command)
			}
		case request := <-transport.Requester:
			live <- 0
			logger.Debug.Printf("request %v", request)
			switch request.Code {
			case transport.CodeReqPhase:
				transport.Sender <- rphase.send()
			case transport.CodeReqPlan:
				transport.Sender <- rplan.send()
			case transport.CodeReqStatus:
				transport.Sender <- rstatus.send()
			case transport.CodeReqSource:
				transport.Sender <- rsource.send()
			case transport.CodeReqSignalGroups:
				transport.Sender <- rsg.send()
			case transport.CodeReqAlarm:
				transport.Sender <- ralarm.send()
			case transport.CodeReqMajor:
				transport.Sender <- rmajor.send()
			default:
				logger.Error.Printf("not request %v", request)
			}
		}
	}
}
