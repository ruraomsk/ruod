package web

import (
	"fmt"
	"time"

	"github.com/anoshenko/rui"
	"github.com/ruraomsk/ag-server/logger"
	"github.com/ruraomsk/ruod/controller/transport"
	"github.com/ruraomsk/ruod/hardware"
)

const controlText = `
		ListLayout {
			style = showPage,
			orientation = vertical,
			content = [
				TextView {
					style=header1,
					text = "<b>Текущее состояние Центра управления</b>",
				},
				ListLayout {
					orientation = horizontal, list-column-gap=16px,padding = 16px,
					border = _{style=solid,width=4px,color=blue },
					content = [
						TextView {
							style=header1,
							id=nowAllRed,text="Кругом Красный"
						},
						TextView {
							style=header1,
							id=nowFlashing,text="Желтое Мигание"
						},
						TextView {
							style=header1,
							id=nowDark,text="Выключено"
						},
						TextView {
							style=header1,
							id=nowPlan,text=""
						},
						TextView {
							style=header1,
							id=nowPhase,text=""
						},
					]
				},
				TextView {
					style=header1,
					text = "<b>Команды от имени центра управления </b>",
				},
				ListLayout {
					orientation = horizontal, list-column-gap=16px,padding = 16px,
					border = _{style=solid,width=4px,color=blue },
					content = [
						Button {
							id=setAllRedOn,content="КК on"
						},
						Button {
							id=setAllRedOff,content="КК off"
						},
						Button {
							id=setFlashingOn,content="ЖМ on"
						},
						Button {
							id=setFlashingOff,content="ЖМ off"
						},
						Button {
							id=setDarkOn,content="ОС on"
						},
						Button {
							id=setDarkOff,content="ОС off"
						},
						Button {
							id=setPlan,content="Установить План"
						},
						NumberPicker {
							id=idPlan,type=editor,min=0,max=32,value=0
						},
						Button {
							id=setPhase,content="Установить Фазу"
						},
						NumberPicker {
							id=idPhase,type=editor,min=0,max=32,value=0
						},
					]
				},
			]
		}
`

func makeViewControl(view rui.View) {
	mutex.Lock()
	defer mutex.Unlock()
	hs := hardware.GetStateHard()
	if hs.Dark {
		rui.Set(view, "nowDark", "visibility", "visible")
	} else {
		rui.Set(view, "nowDark", "visibility", "invisible")
	}
	if hs.AllRed {
		rui.Set(view, "nowAllRed", "visibility", "visible")
	} else {
		rui.Set(view, "nowAllRed", "visibility", "invisible")
	}
	if hs.Flashing {
		rui.Set(view, "nowFlashing", "visibility", "visible")
	} else {
		rui.Set(view, "nowFlashing", "visibility", "invisible")
	}
	rui.Set(view, "nowPlan", "text", fmt.Sprintf("План %d", hs.Plan))
	rui.Set(view, "nowPhase", "text", fmt.Sprintf("Фаза %d", hs.Phase))

}
func updaterControl(view rui.View, session rui.Session) {
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		if view == nil {
			return
		}
		w, ok := SessionStatus[session.ID()]
		if !ok {
			continue
		}

		if !w {
			continue
		}
		makeViewControl(view)
	}

}
func controlKDM(session rui.Session) rui.View {
	view := rui.CreateViewFromText(session, controlText)
	if view == nil {
		return nil
	}
	rui.Set(view, "setAllRedOn", rui.ClickEvent, func(rui.View) {
		transport.FromWeb <- transport.Command{Code: transport.CodeCallAllRed, Value: 1}
		logger.Info.Printf("Оператор установил КК")
	})
	rui.Set(view, "setAllRedOff", rui.ClickEvent, func(rui.View) {
		transport.FromWeb <- transport.Command{Code: transport.CodeCallAllRed, Value: 0}
		logger.Info.Printf("Оператор отменил КК")
	})
	rui.Set(view, "setFlashingOn", rui.ClickEvent, func(rui.View) {
		transport.FromWeb <- transport.Command{Code: transport.CodeCallFlash, Value: 1}
		logger.Info.Printf("Оператор установил ЖМ")
	})
	rui.Set(view, "setFlashingOff", rui.ClickEvent, func(rui.View) {
		transport.FromWeb <- transport.Command{Code: transport.CodeCallFlash, Value: 0}
		logger.Info.Printf("Оператор отменил ЖМ")
	})
	rui.Set(view, "setDarkOn", rui.ClickEvent, func(rui.View) {
		transport.FromWeb <- transport.Command{Code: transport.CodeCallDark, Value: 1}
		logger.Info.Printf("Оператор установил ОС")
	})
	rui.Set(view, "setDarkOff", rui.ClickEvent, func(rui.View) {
		transport.FromWeb <- transport.Command{Code: transport.CodeCallDark, Value: 0}
		logger.Info.Printf("Оператор отменил ОС")
	})
	rui.Set(view, "setPlan", rui.ClickEvent, func(rui.View) {
		transport.FromWeb <- transport.Command{Code: transport.CodeCallPlan, Value: getInteger(rui.Get(view, "idPlan", "value"))}
		logger.Info.Printf("Оператор вызвал план %d", getInteger(rui.Get(view, "idPlan", "value")))
	})
	rui.Set(view, "setPhase", rui.ClickEvent, func(rui.View) {
		transport.FromWeb <- transport.Command{Code: transport.CodeCallPhase, Value: getInteger(rui.Get(view, "idPhase", "value"))}
		logger.Info.Printf("Оператор вызвал фазу %d", getInteger(rui.Get(view, "idPhase", "value")))
	})

	makeViewControl(view)
	go updaterControl(view, session)
	return view
}
