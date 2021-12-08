package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/leoneville/cursogolang/sql/connection"
)

func main() {
	db, err := connection.Conn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")

	stmt.Exec("Steve Rogers", 1)
	stmt.Exec("Tony Stark", 2)

	stmt2, _ := db.Prepare("delete from usuarios where id = ?")
	stmt2.Exec(3)
}
