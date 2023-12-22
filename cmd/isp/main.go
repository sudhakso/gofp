package main

import "fmt"

type Document struct {
	content string
}

// scans & stores the document before printing
type PrintAheadScanner struct {
	mfd MultiFunctionDevice
}

func (pas *PrintAheadScanner) Print(d Document) {
	pas.mfd.Scan(&d)
	pas.mfd.Print(&d)
}

type Printer interface {
	Print(d *Document)
}

type Scanner interface {
	Scan(d *Document)
}

// Interface subclassing
type MultiFunctionDevice interface {
	Printer
	Scanner
}

// Decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

// MFD implementation
type HPEPrinter struct {
}

func (h *HPEPrinter) Print(d *Document) {
	fmt.Println("HPE printing started...")
}

func (h *HPEPrinter) Scan(d *Document) {
	fmt.Println("HPE scanning started...")
}

func main() {
	mfd := PrintAheadScanner{&HPEPrinter{}}
	mfd.Print(Document{"this is Fun!"})
}
