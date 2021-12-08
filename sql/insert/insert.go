package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/leoneville/cursogolang/sql/connection"
)

func main() {
	db, err := connection.Conn()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("insert into usuarios(nome) values(?)")
	stmt.Exec("Maria")
	stmt.Exec("João")

	res, _ := stmt.Exec("Neville")

	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()
	fmt.Println(linhas)
}
