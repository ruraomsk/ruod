package setup

var (
	Set    *Setup
	ExtSet *ExtSetup
)

type Setup struct {
	LogPath string `toml:"logpath"`
	Id      int    `toml:"id"`
	Modbus  Modbus `toml:"modbus" json:"modbus"`
	STCIP   STCIP  `toml:"stcip" json:"stcip"`
}
type ExtSetup struct {
	Modbus Modbus `toml:"modbus" json:"modbus"`
	STCIP  STCIP  `toml:"stcip" json:"stcip"`
}

type Modbus struct {
	Device   string `toml:"device" json:"device"`
	BaudRate int    `toml:"baudrate" json:"baudrate"`
	Parity   string `toml:"parity" json:"parity"`
	UId      int    `toml:"uid" json:"uid"`
	Debug    bool   `toml:"debug" json:"debug"`
	Log      bool   `toml:"log" json:"log"`
}
type STCIP struct {
	Debug  bool   `toml:"debug" json:"debug"`
	Host   string `toml:"host" json:"host"`
	Port   int    `toml:"port" json:"port"`
	Listen int    `toml:"listen" json:"listen"`
}

func (s *Setup) Update(es ExtSetup) {
	s.Modbus = es.Modbus
	s.STCIP = es.STCIP
}
func (es *ExtSetup) Update(s Setup) {
	es.Modbus = s.Modbus
	es.STCIP = s.STCIP
}
