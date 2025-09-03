package rich

import (
	"fmt"
	"io"
	"os"
)

var outWriter io.Writer = os.Stdout

type Styler interface {
	ToANSI() string
}

type Style string

const (
	Reset     Style = "\033[0m"
	Bold      Style = "\033[1m"
	Dim       Style = "\033[2m"
	Italic    Style = "\033[3m"
	Underline Style = "\033[4m"
	Blink     Style = "\033[5m"
	Reverse   Style = "\033[7m"
	Hidden    Style = "\033[8m"
)

func (s Style) ToANSI() string {
	return string(s)
}

func splitArgsAndStyle(argsAndStyles ...any) (args []any, seq string) {
	for _, a := range argsAndStyles {
		if s, ok := a.(Styler); ok {
			seq += s.ToANSI()
		} else {
			args = append(args, a)
		}
	}
	return
}

func formatAndPrint(addNewLine bool, s string, argsAndStyles ...any) string {
	args, seq := splitArgsAndStyle(argsAndStyles...)
	text := fmt.Sprintf(s, args...)
	if addNewLine {
		text += "\n"
	}
	final := seq + text + string(Reset)
	if outWriter != nil {
		_, _ = fmt.Fprint(outWriter, final)
	}
	return final
}

func Println(s string, argsAndStyles ...any) {
	formatAndPrint(true, s, argsAndStyles...)
}

func Printf(s string, argsAndStyles ...any) {
	formatAndPrint(false, s, argsAndStyles...)
}

func Sprint(s string, argsAndStyles ...any) string {
	return formatAndPrint(false, s, argsAndStyles...)
}

func Sprintf(s string, argsAndStyles ...any) string {
	return formatAndPrint(false, s, argsAndStyles...)
}
