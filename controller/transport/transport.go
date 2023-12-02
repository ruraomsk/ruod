package transport

var Commander chan Command
var Requester chan Request
var Sender chan ReplayToServer

const (
	CodeCallPhase = iota
	CodeCallPlan
	CodeCallFlash
	CodeCallAllRed
	CodeCallDark
)
const (
	CodeReqStatus = iota
	CodeReqMajor
	CodeReqPlan
	CodeReqSource
	CodeReqPhase
	CodeReqSignalGroups
	CodeReqAlarm
)

func init() {
	Commander = make(chan Command)
	Requester = make(chan Request)
	Sender = make(chan ReplayToServer)
}

type Command struct {
	OID   string
	Code  int
	Value int
}
type Request struct {
	OID  string
	Code int
}
type ReplayToServer struct {
	Code    int
	Elemets []Element
}
type Element struct {
	Type  int //0 - time 1 - int
	Value uint64
}

type Define struct {
	Code int
	OID  string
}

func Transport() {
	go receiverCommands()
	go receiverRequests()
	go senderReplay()
}
