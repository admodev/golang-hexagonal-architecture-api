package bootstrap

import (
	"bctec/internal/platform/server"
	"bctec/internal/platform/storage/mysql"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// TODO: Move this to separate file...
var host = goDotEnvVariable("HOST")
var portNum = goDotEnvVariable("PORT")
var dbUser = goDotEnvVariable("DB_USER")
var dbPass = goDotEnvVariable("DB_PASS")
var dbHost = goDotEnvVariable("DB_HOST")
var dbPort = goDotEnvVariable("DB_PORT")
var dbName = goDotEnvVariable("DB_NAME")
var port, _ = strconv.ParseUint(portNum, 10, 64)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)

	if err != nil {
		return err
	}

	reportRepository := mysql.NewReportRepository(db)
	srv := server.New(host, uint(port), reportRepository)

	return srv.Run()
}
