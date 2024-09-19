package main

import (
	"database/sql"
	"fmt"

	go_ora "github.com/sijms/go-ora/v2"
)

func main() {
	urlOptions := map[string]string{
		"SID": "xe",
	}
	connStr := go_ora.BuildUrl("localhost", 1521, "", "SYSTEM", "12345", urlOptions)
	conn, err := sql.Open("oracle", connStr)

	// check for error
	if err != nil {
		fmt.Println("Can't open the driver: ", err)
		return
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			fmt.Println("Can't close connection: ", err)
		}
	}()

	err = conn.Ping()
	if err != nil {
		fmt.Println("Can't ping connection: ", err)
		return
	}

	rows, err := conn.Query("SELECT 'Hello World! 123' FROM dual")
	if err != nil {
		fmt.Println("Erro ao fazer select: ", err)
		return
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			fmt.Println("Can't close dataset: ", err)
		}
	}()

	var (
		name string
	)
	for rows.Next() {
		err = rows.Scan(&name)
		fmt.Println("\tName: ", name)
	}
}
