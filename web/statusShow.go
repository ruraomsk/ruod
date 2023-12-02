package web

import (
	"fmt"
	"time"

	"github.com/anoshenko/rui"

	"github.com/ruraomsk/ruod/hardware"
	"github.com/ruraomsk/ruod/setup"
)

// border = _{ style = solid, width = 1px, color = darkgray },
const statusText = `
		ListLayout {
			style = showPage,
			orientation = vertical,
			padding="16px",
			content = [
				TextView {
					style=header1,
					id=titleStatus,text = "",
				},
				TextView {
					id=idUtopia,
					text = "",
					text-size="24px",
				},
				TextView {
					id=idModbus,
					text = "",
					text-size="24px",
				},
			]
		}
`

func toString(t time.Time) string {
	return fmt.Sprintf("%04d/%02d/%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
func makeViewStatus(view rui.View) {
	mutex.Lock()
	defer mutex.Unlock()
	t := time.Now()
	rui.Set(view, "titleStatus", "text", fmt.Sprintf("<b>Текущее состояние контроллера %d </b>%02d:%02d:%02d", setup.Set.Id,
		t.Hour(), t.Minute(), t.Second()))
	c := "Соединение с Центром "
	if hardware.StateHardware.GetConnectCentral() {
		c += "установлено"
	} else {
		c += "отсутствует"
	}
	rui.Set(view, "idUtopia", "text", c)

	c = fmt.Sprintf("Соединение Modbus device %s baud %d parity %s uid %d \t",
		setup.Set.Modbus.Device, setup.Set.Modbus.BaudRate, setup.Set.Modbus.Parity, setup.Set.Modbus.UId)
	if hardware.StateHardware.GetConnect() {
		c += "установлено"
	} else {
		c += "отсутствует"
	}
	rui.Set(view, "idModbus", "text", c)
}
func updaterStatus(view rui.View, session rui.Session) {
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
		makeViewStatus(view)
	}
}

func statusShow(session rui.Session) rui.View {
	view := rui.CreateViewFromText(session, statusText)
	if view == nil {
		return nil
	}
	makeViewStatus(view)
	go updaterStatus(view, session)

	return view
}
