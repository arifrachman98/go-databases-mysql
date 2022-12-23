package godatabases

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Open() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go-databases")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func SetConn() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go-databases")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func TestDatabases(t *testing.T) {

}

func TestOpenConn(t *testing.T) {
	Open()
}
