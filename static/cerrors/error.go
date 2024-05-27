//go:build js && wasm
package cerrors

type InsertErr struct {
    Msg string
}

func (e *InsertErr) Error() string {
    return e.Msg 
}

type CreateErr struct {
    Msg string
}

func (e *CreateErr) Error() string {
    return e.Msg
}

type SelectErr struct {
    Msg string
}

func (e *SelectErr) Error() string {
    return e.Msg
}

type DefaultErr struct {
    Msg string
}

func (e *DefaultErr) Error() string{
    return e.Msg
}
