package main

import (
	"database/sql"

	_ "github.com/cachengo/go-aergolite"
)

func main() {
	for _, driver := range sql.Drivers() {
		println(driver)
	}
}
