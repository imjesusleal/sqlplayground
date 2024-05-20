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

	err = db.Exec("create table whatever(id int, name varchar(25), age int)")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Exec("insert into whatever(id, name, age) values(1, 'jesus', 29)")
	if err != nil {
		log.Fatal(err)
	}

	js.Global().Set("dbConnect", dbConnect(db))

	<-wait

}

func dbConnect(db *sqlite3.Conn) js.Func {
	f := js.FuncOf(func(this js.Value, args []js.Value) any {
		st := args[0].String()
		ret := checkQuery(db, st)
		return ret
	})

	return f

}

func checkQuery(db *sqlite3.Conn, query string) string {
	var st string
	var ret string
	for _, v := range query {
		if string(v) == " " {
			break
		}
		st = st + string(v)
	}
	switch st {
	case "create":
		ret = createTable(db, query)
	case "insert":
		ret = insertQuery(db, query)
	case "select":
		ret = selectQuery(db, query)
	default:
		ret = execQuery(db, query)
	}
	return ret
}

func execQuery(db *sqlite3.Conn, query string) string {
	err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	return "query done sucessfully"
}

func createTable(db *sqlite3.Conn, query string) string {
	err := db.Exec(query)
	if err != nil {
		log.Fatal("Cant execute query, err -> ", err)
	}
	return "1 table succesfully created"
}

func insertQuery(db *sqlite3.Conn, query string) string {
	err := db.Exec(query)
	if err != nil {
		log.Fatal("Cant execute query, err -> ", err)
	}
	return "Insert sucessfully done!"
}

func selectQuery(db *sqlite3.Conn, query string) string {
	var ret string
	stmt, _, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	for stmt.Step() {
		ret = ret + fmt.Sprintln(stmt.ColumnInt(0), stmt.ColumnText(1))
	}
	if err := stmt.Err(); err != nil {
		log.Fatal(err)
	}
	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}
	return ret
}
