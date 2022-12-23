package godatabases

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Open() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go-databases")
	if err != nil {
		panic(err)
	}

	return db

}

func SetConn() *sql.DB {

	dat := Open()

	dat.SetMaxIdleConns(10)
	dat.SetMaxOpenConns(100)
	dat.SetConnMaxIdleTime(5 * time.Minute)
	dat.SetConnMaxLifetime(60 * time.Minute)

	return dat
}

func TestInsertDatabases(t *testing.T) {
	db := SetConn()
	defer db.Close()

	ctx := context.Background()

	masuk := "INSERT INTO customer(id, name) VALUES('2','Jajang')"

	_, err := db.ExecContext(ctx, masuk)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestOpenConn(t *testing.T) {
	Open()
}
