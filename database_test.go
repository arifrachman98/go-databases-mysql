package godatabases

import (
	"context"
	"database/sql"
	"fmt"
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

func TestInsertDatabases(t *testing.T) {
	db := SetConn()
	defer db.Close()

	ctx := context.Background()

	masuk := "INSERT INTO customer(id, name) VALUES('1','Arif')"

	_, err := db.ExecContext(ctx, masuk)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestOpenConn(t *testing.T) {
	Open()
}
