//go:build js && wasm

package main

import (
	"fmt"
	"log"
	"syscall/js"

	"github.com/ncruces/go-sqlite3"
	_ "github.com/ncruces/go-sqlite3/embed"
)

const memory = ":memory:"

func main() {
	wait := make(chan struct{}, 0)

	db, err := sqlite3.Open(memory)
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
		sret = createTable(db, query)
	case "insert":
		sret = insertQuery(db, query)
	case "select":
		ret, err = selectQuery(db, query)
		if err != "" {
            return err,nil
		}
		return "", ret
	default:
		sret = execQuery(db, query)
	}
	return sret, nil
}

func execQuery(db *sqlite3.Conn, query string) string {
	err := db.Exec(query)
	if err != nil {
        return fmt.Sprint("Can't execute query because of err->", err)
	}
	return "query done sucessfully"
}

func createTable(db *sqlite3.Conn, query string) string {
	err := db.Exec(query)
	if err != nil {
        return fmt.Sprint("Can't execute query because of err->", err)
	}
	return "1 table succesfully created"
}

func insertQuery(db *sqlite3.Conn, query string) string {
	err := db.Exec(query)
	if err != nil {
		return fmt.Sprint("Can't execute query because of err->", err)
	}
	return "Insert sucessfully done!"
}

func selectQuery(db *sqlite3.Conn, query string) ([]interface{},string) {
	objects := make([]interface{}, 0)
	stmt, _, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Sprint("Can't execute query because of err->", err)
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
	return objects,""
}
