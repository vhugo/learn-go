// Requirements and acceptance criteria
//
// we will have two objects (Packt and Normal printer) and two implementations
// (PrinterImpl1 and PrinterImpl2) that we will join by using the Bridge design
// pattern. More or less, we will have the following requirements and acceptance
// criteria:
//
// - A PrinterAPI that accepts a message to print
//
// - An implementation of the API that simply prints the message to the console
// An implementation of the API that prints to an io.Writer interface
//
// - A Printer abstraction with a Print method to implement in printing types
//
// A normal printer object, which will implement the Printer and the PrinterAPI
// interface
//
// The normal printer will forward the message directly to the implementation A
// Packt printer, which will implement the Printer abstraction and the
//
// PrinterAPI interface
//
// The Packt printer will append the message Message from Packt: to all prints
//
package main

import "testing"

func TestPrintAPI1(t *testing.T) {
	api1 := PrinterImpl1{}

	err := api1.PrintMessage("Hello")
	if err != nil {
		t.Error("Error trying to use the API1 implementation: Message: %s\n",
			err.Error())
	}
}

func TestPrintAPI2(t *testing.T) {
	testWriter := TestWriter{}
	api2 := PrinterImpl2{
		Writer: &testWriter,
	}

	expectedMessage := "Hello"
	err := api2.PrintMessage(expectedMessage)
	if err != nil {
		t.Errorf("Error trying to use the API2 implementation: %s\n", err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Fatalf("API2 did not write correctly on the io.Writer. \n  Actual: %s\n"+
			"Expected: %s\n", testWriter.Msg, expectedMessage)
	}
}

func TestNormalPrinter_Print(t *testing.T) {
	expectedMessage := "Hello io.Writer"

	normal := NormalPrinter{
		Msg:     expectedMessage,
		Printer: &PrinterImpl1{},
	}

	err := normal.Print()
	if err != nil {
		t.Errorf(err.Error())
	}

	testWriter := TestWriter{}
	normal = NormalPrinter{
		Msg: expectedMessage,
		Printer: &PrinterImpl2{
			Writer: &testWriter,
		},
	}

	err = normal.Print()
	if err != nil {
		t.Error(err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Errorf("The expected message on the io.Writer doesn't match actual.\n"+
			"Actual: %s\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}
}

func TestPacktPrinter_Print(t *testing.T) {
	passedMessage := "Hello io.Writer"
	expectedMessage := "Message from Packt: " + passedMessage

	packt := PacktPrinter{
		Msg:     passedMessage,
		Printer: &PrinterImpl1{},
	}

	err := packt.Print()
	if err != nil {
		t.Errorf(err.Error())
	}

	testWriter := TestWriter{}
	packt = PacktPrinter{
		Msg: passedMessage,
		Printer: &PrinterImpl2{
			Writer: &testWriter,
		},
	}

	err = packt.Print()
	if err != nil {
		t.Errorf(err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Errorf("The expected message on the io.Writer doesn't match actual.\n"+
			"Actual: %s\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}
}
