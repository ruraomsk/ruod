package transport

var comOids = []Define{
	{Code: CodeCallPhase, OID: "1.3.6.1.4.1.1618.3.7.2.11.1"},
	{Code: CodeCallPlan, OID: "1.3.6.1.4.1.1618.3.7.2.2.1"},
	{Code: CodeCallFlash, OID: "1.3.6.1.4.1.1618.3.2.2.1.1"},
	{Code: CodeCallDark, OID: "1.3.6.1.4.1.1618.3.2.2.2.1"},
}

func receiverCommands() {

	for {
		code := <-FromWeb
		for _, v := range comOids {
			if code.Code == v.Code {
				Commander <- Command{OID: v.OID, Code: v.Code, Value: code.Value}
				break
			}
		}
	}
}
