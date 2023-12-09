package setup

var (
	Set    *Setup
	ExtSet *ExtSetup
)

type Setup struct {
	LogPath     string      `toml:"logpath"`
	Id          int         `toml:"id"`
	Modbus      Modbus      `toml:"modbus" json:"modbus"`
	STCIP       STCIP       `toml:"stcip" json:"stcip"`
	TrafficData TrafficData `toml:"trafficdata" json:"trafficdata"`
	ModbusRadar ModbusRadar `toml:"modbusradar" json:"modbusradar"`
}
type ExtSetup struct {
	Modbus      Modbus      `toml:"modbus" json:"modbus"`
	STCIP       STCIP       `toml:"stcip" json:"stcip"`
	TrafficData TrafficData `toml:"trafficdata" json:"trafficdata"`
	ModbusRadar ModbusRadar `toml:"modbusradar" json:"modbusradar"`
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
type ModbusRadar struct {
	Radar   bool   `toml:"radar" json:"radar"`
	Master  bool   `toml:"master" json:"master"`
	Debug   bool   `toml:"debug" json:"debug"`
	Host    string `toml:"host" json:"host"`
	Port    int    `toml:"port" json:"port"`
	ID      int    `toml:"id" json:"id"`
	Chanels int    `toml:"chanels" json:"chanels"`
	Diaps   int    `toml:"diaps" json:"diaps"`
	Diap    int    `toml:"diap" json:"diap"`
}
type TrafficData struct {
	Work    bool   `toml:"work" json:"work"`
	Debug   bool   `toml:"debug" json:"debug"`
	Host    string `toml:"host" json:"host"`
	Port    int    `toml:"port" json:"port"`
	Listen  int    `toml:"listen" json:"listen"`
	Chanels int    `toml:"chanels" json:"chanels"`
	Diaps   int    `toml:"diaps" json:"diaps"`
	Diap    int    `toml:"diap" json:"diap"`
}

func (s *Setup) Update(es ExtSetup) {
	s.Modbus = es.Modbus
	s.STCIP = es.STCIP
	s.ModbusRadar = es.ModbusRadar
	s.TrafficData = es.TrafficData
}
func (es *ExtSetup) Update(s Setup) {
	es.Modbus = s.Modbus
	es.STCIP = s.STCIP
	es.ModbusRadar = s.ModbusRadar
	es.TrafficData = s.TrafficData
}
