package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/leoneville/cursogolang/sql/connection"
)

func main() {
	db, err := connection.Conn("mysql", "root:123@123@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")

	stmt.Exec(4000, "Bia")
	stmt.Exec(4001, "Carlos")
	_, err = stmt.Exec(1, "Thiago") // Chave duplicada

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
