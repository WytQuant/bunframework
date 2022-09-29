package connectdb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"log"
	"os"
)

type DbInstance struct {
	Db *bun.DB
}

var Database DbInstance

func Connect() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Cannot load .env file")
	}

	dsn := os.Getenv("DSN")
	sqldb, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
		return
	}

	bundb := bun.NewDB(sqldb, mysqldialect.New())

	Database = DbInstance{Db: bundb}
}
