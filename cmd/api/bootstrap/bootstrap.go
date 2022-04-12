package bootstrap

import (
	"bctec/internal/platform/server"
	"bctec/internal/platform/storage/mysql"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host   = "localhost"
	port   = 8080
	dbUser = "root"
	dbPass = "admocode"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "bctec"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)

	if err != nil {
		return err
	}

	reportRepository := mysql.NewReportRepository(db)
	srv := server.New(host, port, reportRepository)

	return srv.Run()
}
