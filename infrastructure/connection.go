/**
* Created by Maulana Ibrahim
* on 27 June 2020
 */

package infrastructure

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func InitDB(driver, dataSource string) (*sql.DB, error) {
	db, _ := sql.Open(driver, dataSource)

	if err := db.Ping(); err != nil {
		return db, err
	}
	return db, nil
}

func ConnectDB() *sql.DB {
	env := Environment()
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", env.dbUser, env.dbPassword, env.dbHost, env.dbPort, env.schemaName)
	db, err := InitDB("mysql", dataSource)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
