package environment

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

var HOST = GoDotEnvVariable("HOST")
var PORTNUM = GoDotEnvVariable("PORT")
var DBUSER = GoDotEnvVariable("DB_USER")
var DBPASS = GoDotEnvVariable("DB_PASS")
var DBHOST = GoDotEnvVariable("DB_HOST")
var DBPORT = GoDotEnvVariable("DB_PORT")
var DBNAME = GoDotEnvVariable("DB_NAME")
var PORT, _ = strconv.ParseUint(PORTNUM, 10, 64)
var SECRET = GoDotEnvVariable("SECRET")
