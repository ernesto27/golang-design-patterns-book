package adapter

import "testing"

func TestAdap(t *testing.T) {
	msg := "Hello World!"

	adapter := PrinterAdapter{OldPrinter: &MyLegacyPriner{}, Msg: msg}

	returnedMsg := adapter.PrintStored()

	if returnedMsg != "Legacy Printer: Adapter: Hello World!\n" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}

	adapterNil := PrinterAdapter{OldPrinter: nil, Msg: msg}
	returnedMsg2 := adapterNil.PrintStored()

	if returnedMsg2 != "Hello World!" {
		t.Errorf("Message didn't match: %s\n", returnedMsg2)
	}

}
