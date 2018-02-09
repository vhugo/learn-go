// Adapter
//
// Like in real life, where you have plug adapters and bolt adapters, in Go, an
// adapter will allow us to use something that wasn't built for a specific task
// at the beginning.
//
// Adapter also helps us to maintain the open/closed principle in our apps,
// making them more predictable too. They also allow us to write code which uses
// some base that we can't modify.
//
package main

import "fmt"

type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyPrinter struct{}

func (l *MyLegacyPrinter) Print(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy Printer: %s\n", s)
	println(newMsg)
	return
}

type ModernPrinter interface {
	PrintStored() string
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *PrinterAdapter) PrintStored() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)

	} else {
		newMsg = p.Msg
	}
	return
}
