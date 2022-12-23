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

func OpenConn() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/godb?parseTime=true")
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

func SetConnect() *sql.DB {

	dat := OpenConn()

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

func TestQuerySQLComplex(t *testing.T) {
	db := SetConnect()
	defer db.Close()

	ctx := context.Background()
	perintah := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM pelanggan"
	rows, err := db.QueryContext(ctx, perintah)

	if err != nil {
		panic(err)
	}

	//retrive data from databases
	for rows.Next() {
		var (
			name        string
			email       sql.NullString
			id, balance int32
			rating      float32
			birthdate   sql.NullTime
			createdAt   time.Time
			married     bool
		)
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthdate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("====================")
		fmt.Println("ID :", id)
		fmt.Println("Nama :", name)
		if email.Valid {
			fmt.Println("Email :", email.String)
		}
		fmt.Println("Balance :", balance)
		fmt.Println("Rating :", rating)
		if birthdate.Valid {
			fmt.Println("Birth Date :", birthdate.Time)
		}
		fmt.Println("Married :", married)
		fmt.Println("Created At :", createdAt)
	}

	fmt.Println("Success execute query table")
	defer rows.Close()
}

func TestSQLInjection(t *testing.T) {
	db := SetConnect()
	defer db.Close()
	ctx := context.Background()

	uname := "admin"
	passw := "admin"

	perintah := "SELECT username FROM user WHERE username = '" + uname + "' AND password = '" + passw + "' LIMIT 1"

	rows, err := db.QueryContext(ctx, perintah)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var uname string

		err := rows.Scan(&uname)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login, Welcome", uname)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestSQLInjectionSec(t *testing.T) {
	db := SetConnect()
	defer db.Close()
	ctx := context.Background()

	uname := "admin';#"
	passw := "salah"

	perintah := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, perintah, uname, passw)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var uname string

		err := rows.Scan(&uname)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login, Welcome", uname)
	} else {
		fmt.Println("Gagal Login")
	}
}
