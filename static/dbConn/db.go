//go:build js && wasm

package dbConn

import (

	"github.com/ncruces/go-sqlite3"
	_ "github.com/ncruces/go-sqlite3/embed"
)

const memory = ":memory:"

type IDbConn interface {
    Open() (*sqlite3.Conn, error)
    Close()
}

type DbConn struct {
    db *sqlite3.Conn
}

func (d *DbConn) Open() (*sqlite3.Conn,error) {
    var err error
    d.db, err = sqlite3.Open(memory)
    if err != nil {
        return nil, err
    }
    return d.db, nil 
}

func (d *DbConn) Close() {
    if d.db != nil {
        d.db.Close()
    }
}

func CreateDb() IDbConn {
    return &DbConn{}
}
