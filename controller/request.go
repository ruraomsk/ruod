package controller

import (
	"fmt"
	"time"

	"github.com/ruraomsk/ruod/controller/transport"
	"github.com/ruraomsk/ruod/hardware"
)

// Запросы центра
// статус - 1.3.6.1.4.1.1618.3.6.2.1.2.0
// code=1
type RepStatus struct {
	Time   time.Time
	Status int
}

func (r *RepStatus) send() transport.ReplayToServer {
	r.Time = time.Now()
	s := hardware.GetStateHard()
	r.Status = int(s.Status[0])
	return transport.ReplayToServer{Code: transport.CodeReqStatus, Elemets: make([]transport.Element, 0)}
}
func (r *RepStatus) ToSting() string {
	return fmt.Sprintf("%s Статус %d", timeToString(r.Time), r.Status)
}

// режим работы - 1.3.6.1.4.1.1618.3.6.2.2.2.0
// SwarcoUTCStatusMode OBJECT-TYPE
// SYNTAX OperationalMode
// ACCESS read-only
// STATUS current
// DESCRIPTION
// "The major operational mode of the traffic
// controller."
// ::= { SwarcoUTCStatusModeObjs 2 }
// code=2
type RepMajor struct {
	Time   time.Time
	Status int
}

func (r *RepMajor) send() transport.ReplayToServer {
	r.Time = time.Now()
	s := hardware.GetStateHard()
	r.Status = int(s.Status[0])
	return transport.ReplayToServer{Code: transport.CodeReqMajor, Elemets: make([]transport.Element, 0)}
}
func (r *RepMajor) ToSting() string {
	return fmt.Sprintf("%s Состояние %d", timeToString(r.Time), r.Status)
}

// номер плана - 1.3.6.1.4.1.1618.3.5.2.1.7.0
// SwarcoUTCSignalGroupPlanNo OBJECT-TYPE
// SYNTAX Unsigned32
// ACCESS read-only
// STATUS current
// DESCRIPTION
// "Plan number"
// ::= { SwarcoUTCSignalGroupStateObjs 7 }
// code=3
type RepPlan struct {
	Time time.Time
	Plan int
}

func (r *RepPlan) send() transport.ReplayToServer {
	r.Time = time.Now()
	r.Plan = hardware.GetPlan()
	return transport.ReplayToServer{Code: transport.CodeReqPlan, Elemets: make([]transport.Element, 0)}
}
func (r *RepPlan) ToSting() string {
	return fmt.Sprintf("%s План координации %d", timeToString(r.Time), r.Plan)
}

// источник плана - 1.3.6.1.4.1.1618.3.7.2.1.3.0
// SwarcoUTCTrafftechPlanSource OBJECT-TYPE
// SYNTAX PlanSource
// ACCESS read-only
// STATUS current
// DESCRIPTION
// "The reason for the current traffic plan."
// ::= { SwarcoUTCTrafftechPlanStatusObjs 3 }
// code=4
type RepSource struct {
	Time   time.Time
	Source int
}

func (r *RepSource) send() transport.ReplayToServer {
	r.Time = time.Now()
	return transport.ReplayToServer{Code: transport.CodeReqSource, Elemets: make([]transport.Element, 0)}
}
func (r *RepSource) ToSting() string {
	return fmt.Sprintf("%s Источник %d", timeToString(r.Time), r.Source)
}

// номер фазы - 1.3.6.1.4.1.1618.3.7.2.11.2.0
// SwarcoUTCTrafftechPhaseStatus OBJECT-TYPE
// SYNTAX Unsigned32
// ACCESS read-only
// STATUS current
// DESCRIPTION
// "Current phase."
// ::= { SwarcoUTCTrafftechPhaseObjs 2 }
// code=5
type RepPhase struct {
	Time  time.Time
	Phase int
}

func (r *RepPhase) send() transport.ReplayToServer {
	r.Time = time.Now()
	r.Phase = hardware.GetPhase()
	return transport.ReplayToServer{Code: transport.CodeReqPhase, Elemets: make([]transport.Element, 0)}
}
func (r *RepPhase) ToSting() string {
	return fmt.Sprintf("%s Запрос фазы %d", timeToString(r.Time), r.Phase)
}

// состояние сигнальных групп - 1.3.6.1.4.1.1618.3.5.2.1.6.0
// SwarcoUTCSignalGroupState OBJECT-TYPE
// SYNTAX SignalGroupState
// ACCESS read-only
// STATUS current
// DESCRIPTION
// "The current state of the signal groups."
// ::= { SwarcoUTCSignalGroupStateObjs 6 }
// code=6
type RepSignalGroups struct {
	Time           time.Time // Time, when signal group status was changed
	CycleTime      int       //  Cycle Time, program cycle counter
	OffsetTime     int       // Offset Time, program cycle counter with offset if current operation mode is coordinated, stage counter otherwise
	SequenceNumber int       // Sequence Number, packet sequence number used to detect packets arriving out of sequence
	Quantity       int       // Quantity, number of signal groups in the controller.
	State          int       // State, the states of the signal groups.
	Plan           int       // Plan number, current time plan.
	Status         int       // Status, signal group status.
	Plus           int       // Plus all the detector status.
}

func (r *RepSignalGroups) send() transport.ReplayToServer {
	r.Time = time.Now()
	return transport.ReplayToServer{Code: transport.CodeReqSignalGroups, Elemets: make([]transport.Element, 0)}
}
func (r *RepSignalGroups) ToSting() string {
	return fmt.Sprintf("%s Сигнальные группы ", timeToString(r.Time))
}

// тревоги - 1.3.6.1.4.1.1618.3.1.2.2.2.0
// code=7
type RepAlarm struct {
	Time   time.Time
	Count  int //Number of active alarms sent in message (zero to ten, zero if no active alarm exists)
	Alarms []Alarm
}
type Alarm struct {
	Type  int
	Info1 string
	Info  string
}

func (r *RepAlarm) send() transport.ReplayToServer {
	r.Time = time.Now()
	return transport.ReplayToServer{Code: transport.CodeReqAlarm, Elemets: make([]transport.Element, 0)}
}
func (r *RepAlarm) ToSting() string {
	return fmt.Sprintf("%s Тревоги ", timeToString(r.Time))
}
