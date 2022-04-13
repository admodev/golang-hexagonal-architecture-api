package bootstrap

import (
	"bctec/cmd/api/environment"
	"bctec/internal/platform/server"
	"bctec/internal/platform/storage/mysql"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", environment.DBUSER, environment.DBPASS, environment.DBHOST, environment.DBPORT, environment.DBNAME)
	db, err := sql.Open("mysql", mysqlURI)

	if err != nil {
		return err
	}

	reportRepository := mysql.NewReportRepository(db)
	userRepository := mysql.NewUserRepository(db)
	srv := server.New(environment.HOST, uint(environment.PORT), reportRepository, userRepository)

	return srv.Run()
}
