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

func TestOpenConn(t *testing.T) {
	Open()
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

func TestQuerySQL(t *testing.T) {
	db := SetConn()  //initiate connection
	defer db.Close() //close connection at last execution function

	ctx := context.Background() //initiate background context

	perintah := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, perintah)
	if err != nil {
		panic(err)
	}

	//to retrive data from databases
	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("ID :", id)
		fmt.Println("Name :", name)
	}

	fmt.Println("Succes execute query table")

	defer rows.Close()
}
