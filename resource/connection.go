// Package resource provides connection setup utilities, such as database connections.
package resource

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"itsm-itom-ai-assistant/config"
	"log"
	"sync"
)

var once sync.Once
var dbConn *sql.DB

func Connect() *sql.DB {
	once.Do(func() {
		Conf, err := config.LoadConfig()
		if err != nil {
			log.Fatal(err)
		}
		dbConn, err = sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", Conf.DBUser, Conf.DBPassword, Conf.DBName))
		if err != nil {
			log.Fatal(err)
		}
		err = dbConn.Ping()
		log.Println("Successfully connected!")
		if err != nil {
			log.Fatal(err)
		}
	})
	return dbConn
}
