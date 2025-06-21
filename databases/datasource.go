package databases

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/RajaSunrise/simple-bank/utils"
)

var DB *sql.DB

func init() {
	connect := fmt.Sprintf(
		"host=%s, port=%d, user=%s, password=%s, dbname=%s, sslmode=disable, timexone=%s",
		utils.DBHost, utils.DBPort, utils.DBUser, utils.DBPassword, utils.DBName, utils.TimeZone,
	)

	db, err := sql.Open("postgres", connect)
	if err != nil {
		log.Println("Error Connect To databases")
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Println("Error Check Ping databases")
		log.Fatal(err)
	}

	DB = db

	log.Println("Connect To Database", utils.DBName)

}
