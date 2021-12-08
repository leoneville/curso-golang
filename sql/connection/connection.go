package connection

import "database/sql"

// Conn cria a conexão com o banco de dados e retorna um pointer para sql.DB e um error
func Conn(driver, data string) (connection *sql.DB, err error) {
	connection, err = sql.Open(driver, data)
	if err != nil {
		return nil, err
	}
	return
}
