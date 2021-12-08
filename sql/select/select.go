package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/leoneville/cursogolang/sql/connection"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := connection.Conn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select id, nome from usuarios where id > ?", 0)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user usuario
		rows.Scan(&user.id, &user.nome)
		fmt.Println(user)
	}
}
