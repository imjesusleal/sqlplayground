//go:build js && wasm

package cerrors

import (
	"fmt"
	"strings"

	"github.com/ncruces/go-sqlite3"
)

type DefaultErr struct {
	msg string
}

func (e *DefaultErr) Error() string {
	return fmt.Sprintf("Tienes un error de sintaxis cerca de: %s", e.msg)
}

type LogicErr struct {
	msg string
}

func (e *LogicErr) Error() string {
	s := strings.Split(e.msg, " ")
	s = s[4:]
	switch s[0] {
	case "table":
		e.msg = fmt.Sprintf("Tabla %s no tiene una columna llamada %s", s[1], s[len(s)-1])
	case "no":
		e.msg = fmt.Sprintf("No existe la tabla %s", s[3])
	case "incomplete":
		e.msg = fmt.Sprintf("Estas intentando enviar una query incompleta.")
	default:
		e.msg = fmt.Sprintf("%s no tiene %s sobre %s", s[4], s[5], s[8])
	}
	return e.msg

}

func Wrap(err error) string {
	if errs, ok := err.(*sqlite3.Error); ok {
		if len(errs.SQL()) == 0 {
			s := LogicErr{msg: errs.Error()}
			return s.Error()
		}
		e := &DefaultErr{msg: errs.SQL()}
		return e.Error()
	}
	return ""
}
