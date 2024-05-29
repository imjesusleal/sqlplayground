//go:build js && wasm

package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ncruces/go-sqlite3"
	_ "github.com/ncruces/go-sqlite3/embed"
)

var mockDb *sqlite3.Conn
var err error

func init() {
	mockDb, err = sqlite3.Open(memory)
	if err != nil {
		fmt.Println(err)
	}
}

func TestSelectQuery(t *testing.T) {
	err = mockDb.Exec("Create table user(id INT, name VARCHAR(25), age int)")
	if err != nil {
		fmt.Println(err)
	}

	err = mockDb.Exec("Insert into user(id, name, age) values(1, 'jesus', 29)")
	if err != nil {
		fmt.Println(err)
	}

	type MockTests[T any] struct {
		name  string
		DB    *sqlite3.Conn
		query string
		want  T
	}

	tests := []MockTests[[]interface{}]{
		{"Debe devolver todo los campos que encuentre", mockDb, "select * from user", []interface{}{
			map[string]interface{}{"id": int64(1), "name": "jesus", "age": int64(29)}}},
		{"Debe devolver solo los campos: id y name", mockDb, "select id,name from user", []interface{}{
			map[string]interface{}{"id": int64(1), "name": "jesus"}}},
		{"Debe devolver solo el campo id", mockDb, "select id from user", []interface{}{
			map[string]interface{}{"id": int64(1)}}},
		{"Debe devolver el campo age", mockDb, "select age from user", []interface{}{
			map[string]interface{}{"age": int64(29)}}},
		{"Debe devolver todos los campos ordenados por id", mockDb, "select * from user order by id", []interface{}{
			map[string]interface{}{"id": int64(1), "name": "jesus", "age": int64(29)}}},
	}

	testsErr := []MockTests[string]{
		{"Debe devolver una cadena indicando el error de seleccion de tabla", mockDb, "select * from users order by id", "No existe la tabla users"},
		{"Debe devolver una cadena indicando el error de sintaxis en la seleccion de columna", mockDb, "select (ida,name) from user", "Tienes un error de sintaxis cerca de: ida,name) from user"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, _ := SelectQuery(tt.DB, tt.query)

			if !(reflect.DeepEqual(ans, tt.want)) {
				t.Errorf("got %s, want %s", ans, tt.want)
			}

		})
	}

	for _, ff := range testsErr {
		t.Run(ff.name, func(t *testing.T) {
			_, ans := SelectQuery(ff.DB, ff.query)

			if ans != ff.want {
				t.Errorf("got %s, want %s", ans, ff.want)
			}
		})
	}
}

func TestInsertQuery(t *testing.T) {
	err = mockDb.Exec("Create table user(id INT, name VARCHAR(25), age int)")
	if err != nil {
		fmt.Println(err)
	}

	var MockInserTest = []struct {
		name  string
		DB    *sqlite3.Conn
		query string
		want  string
	}{
		{"Debe insertar correctamente la fila", mockDb, "insert into user(id, name) values(2, 'clara')", "La inserción en la tabla se ha hecho correctamente."},
		{"Debe devolver un error", mockDb, "insert into users(id,name) values(2,'clara')", "No existe la tabla users"},
		{"Debe devolver error", mockDb, "", "Estas enviando una consulta vacia"},
	}

	for _, tt := range MockInserTest {
		t.Run(tt.name, func(t *testing.T) {
			ans := InsertQuery(tt.DB, tt.query)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}

}

func TestCreateTable(t *testing.T) {
	type MockCreateTable struct {
		name  string
		DB    *sqlite3.Conn
		query string
		want  string
	}

	tests := []MockCreateTable{
		{"Debe devolver una cadena", mockDb, "Create table if not exists mockUser(id INT PRIMARY KEY, name VARCHAR(25), age INT, is_active boolean null check (is_active in (0,1)))",
			"Se ha creado la tabla correctamente."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := CreateTable(tt.DB, tt.query)

			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
