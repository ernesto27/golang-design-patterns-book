package adapter

import (
	"fmt"
	"io"
	"strconv"
)

type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyPriner struct{}

func (l *MyLegacyPriner) Print(s string) (newMsg string) {
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

func (p *PrinterAdapter) PrintStored() string {
	newMsg := ""
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
	} else {
		newMsg = p.Msg
	}

	return newMsg
}

type Counter struct {
	Writer io.Writer
}

func (f *Counter) Count(n uint64) uint64 {
	if n == 0 {
		f.Writer.Write([]byte(strconv.Itoa(0) + "	\n"))
		return 0
	}

	cur := n
	f.Writer.Write([]byte(strconv.FormatUint(cur, 10) + "\n"))
	return f.Count(n - 1)
}
