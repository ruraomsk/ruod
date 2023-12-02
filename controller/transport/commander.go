package transport

import "time"

var comOids = []Define{
	{Code: CodeCallPhase, OID: "1.3.6.1.4.1.1618.3.7.2.11.1"},
	{Code: CodeCallPlan, OID: "1.3.6.1.4.1.1618.3.7.2.2.1"},
	{Code: CodeCallFlash, OID: "1.3.6.1.4.1.1618.3.2.2.1.1"},
	{Code: CodeCallDark, OID: "1.3.6.1.4.1.1618.3.2.2.2.1"},
}

func receiverCommands() {
	for {
		for _, v := range comOids {
			time.Sleep(time.Second)
			Commander <- Command{OID: v.OID, Code: v.Code, Value: 1}
		}
	}
}
