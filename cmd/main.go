package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/sakhimpungose/ecom/cmd/api"
	"github.com/sakhimpungose/ecom/config"
	"github.com/sakhimpungose/ecom/db"
)

func main() {

	db, err := db.NewMySQLStorage(mysql.Config {
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)
	server := api.NewAPIServer(":"+config.Envs.Port, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Database successfully connected.")
}