// Requirements and acceptance criteria
//
// Having an old interface called LegacyPrinter and a new one called
// ModernPrinter, create a structure that implements the ModernPrinter interface
// and can use the LegacyPrinter interface as described in the following steps:
//
// 1. Create an Adapter object that implements the ModernPrinter interface.
//
// 2. The new Adapter object must contain an instance of the LegacyPrinter
// interface.
//
// 3. When using ModernPrinter, it must call the LegacyPrinter interface under
// the hood, prefixing it with the text Adapter.
//
package main

import "testing"

func TestAdapter(t *testing.T) {
	msg := "Hello World!"
	adapter := PrinterAdapter{OldPrinter: &MyLegacyPrinter{}, Msg: msg}

	returnedMsg := adapter.PrintStored()

	if returnedMsg != "Legacy Printer: Adapter: Hello World!\n" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}

	adapter = PrinterAdapter{OldPrinter: nil, Msg: msg}
	returnedMsg = adapter.PrintStored()

	if returnedMsg != "Hello World!" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}
}
