//go:build js && wasm

package main

import (
	"fmt"
	"log"
	"strings"
	"syscall/js"
	"wasi/static/cerrors"
	"wasi/static/dbConn"

	"github.com/ncruces/go-sqlite3"
	_ "github.com/ncruces/go-sqlite3/embed"
)

func main() {
	wait := make(chan struct{}, 0)
    
    conn := dbConn.CreateDb()
    db, err := conn.Open()
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()

	err = db.Exec("create table user(id int, name varchar(25))")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Exec("insert into user(id, name) values(1, 'jesus')")
	if err != nil {
		log.Fatal(err)
	}

	js.Global().Set("dbConnect", dbConnect(db))

	<-wait

}

func dbConnect(db *sqlite3.Conn) js.Func {
	f := js.FuncOf(func(this js.Value, args []js.Value) any {
		st := args[0].String()
		st = strings.ToLower(st)
		str, ret := checkQuery(db, st)
		if str == "" {
			return ret
		}
		return str
	})

	return f

}

func checkQuery(db *sqlite3.Conn, query string) (string, []interface{}) {
	var st string
	var sret string
	var ret []interface{}
	var err string
	for _, v := range query {
		if string(v) == " " {
			break
		}
		st = st + string(v)
	}
	switch st {
	case "create":
		sret = CreateTable(db, query)
	case "insert":
		sret = InsertQuery(db, query)
	case "select":
		ret, err = SelectQuery(db, query)
		if err != "" {
			return err, nil
		}
		return "", ret
	default:
		sret = execQuery(db, query)
	}
	return sret, nil
}

func execQuery(db *sqlite3.Conn, query string) string {
	if len(query) == 0 {
		return fmt.Sprint("Estas enviando una consulta vacia")
	}
	err := db.Exec(query)
	if err != nil {
		e := cerrors.Wrap(err)
		return e
	}
	return "Query realizada correctamente."
}

func CreateTable(db *sqlite3.Conn, query string) string {
	err := db.Exec(query)
	if err != nil {
		e := cerrors.Wrap(err)
		return e
	}
	return "Se ha creado la tabla correctamente."
}

func InsertQuery(db *sqlite3.Conn, query string) string {
	if len(query) == 0 {
		return fmt.Sprint("Estas enviando una consulta vacia")
	}
	err := db.Exec(query)
	if err != nil {
		e := cerrors.Wrap(err)
		return e
	}
	return "La inserción en la tabla se ha hecho correctamente."
}

func SelectQuery(db *sqlite3.Conn, query string) ([]interface{}, string) {
	objects := make([]interface{}, 0)
	stmt, _, err := db.Prepare(query)
	if err != nil {
		e := cerrors.Wrap(err)
		return nil, e
	}

	size := stmt.ColumnCount()
	rows := make([]interface{}, size)
	for stmt.Step() {
		qmap := make(map[string]interface{}, 0)
		stmt.Columns(rows)
		for i := 0; i < size; i++ {
			qmap[stmt.ColumnName(i)] = rows[i]
		}
		objects = append(objects, qmap)
	}

	if err := stmt.Err(); err != nil {
		log.Fatal(err)
	}
	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}
	return objects, ""
}
