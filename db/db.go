package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/MufidJamaluddin/simplewebgo/config"
	"net/url"
	"log"
	"fmt"
)

// UDAH DI DESAIN "LONG LIVE!" DAN THREAD-SAFE!S 
var dbConnection *sql.DB

// GetConnection : Mendapatkan Koneksi
func GetConnection() *sql.DB {
	return dbConnection
}

func init(){
	connect()
}

func connect(){
	var err error

	configo := config.GetConfig()

	connectionString := configo.Database.GetConnectionString()
	
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", configo.Database.Timezone)

	dsn := fmt.Sprintf("%s?%s", connectionString, val.Encode())
	
	dbConnection, err = sql.Open(configo.Database.Engine, dsn)
	
	if err != nil {
		panic(err)
	}

	err = dbConnection.Ping()
	
	if err != nil {
		panic(err)
	}

	defer func() {
		err := dbConnection.Close()
		if err != nil {
			panic(err)
		}
	}()
}