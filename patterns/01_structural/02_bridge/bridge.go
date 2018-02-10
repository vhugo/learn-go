// Bridge
//
// The Bridge pattern is a design with a slightly cryptic definition from the
// original Gang of Four book. It decouples an abstraction from its
// implementation so that the two can vary independently. This cryptic
// explanation just means that you could even decouple the most basic form of
// functionality: decouple an object from what it does.
//
// It decouples abstraction (an object) from its implementation (the thing that
// the object does). This way, we can change what an object does as much as we
// want. It also allows us to change the abstracted object while reusing the
// same implementation.
//
// The objective of the Bridge pattern is to bring flexibility to a struct that
// change often. Knowing the inputs and outputs of a method, it allows us to
// change code without knowing too much about it and leaving the freedom for
// both sides to be modified more easily.
//
package main

import (
	"errors"
	"fmt"
	"io"
)

type PrinterAPI interface {
	PrintMessage(string) error
}

type PrinterImpl1 struct{}

func (p *PrinterImpl1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg)
	return nil
}

type PrinterImpl2 struct {
	Writer io.Writer
}

func (p *PrinterImpl2) PrintMessage(msg string) error {
	if p.Writer == nil {
		return errors.New("You need to pass an io.Writer to PrinterImpl2")
	}

	fmt.Fprintf(p.Writer, "%s", msg)
	return nil
}

type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}
	err = errors.New("Content received on Writer was empty")
	return
}

type PrinterAbstraction interface {
	Print() error
}

type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (c *NormalPrinter) Print() error {
	c.Printer.PrintMessage(c.Msg)
	return nil
}

type PacktPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (c *PacktPrinter) Print() error {
	c.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %s", c.Msg))
	return nil
}
