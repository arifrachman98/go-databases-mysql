package godatabases

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//DB go-databases
func Open() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go-databases")
	if err != nil {
		panic(err)
	}

	return db

}

//DB godb
func OpenConn() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/godb?parseTime=true")
	if err != nil {
		panic(err)
	}

	return db
}

//DB go-databases
func SetConn() *sql.DB {

	dat := Open()

	dat.SetMaxIdleConns(10)
	dat.SetMaxOpenConns(100)
	dat.SetConnMaxIdleTime(5 * time.Minute)
	dat.SetConnMaxLifetime(60 * time.Minute)

	return dat
}

//DB godb
func SetConnect() *sql.DB {

	dat := OpenConn()

	dat.SetMaxIdleConns(10)
	dat.SetMaxOpenConns(100)
	dat.SetConnMaxIdleTime(5 * time.Minute)
	dat.SetConnMaxLifetime(60 * time.Minute)

	return dat
}

//DB go-databases/customer
func TestOpenConn(t *testing.T) {
	Open()
}

//DB godb/pelanggan
func TestOpenConnSec(t *testing.T) {
	OpenConn()
}

//DB go-databases/customer
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

//DB godb/pelanggan
func TestInsertDataSec(t *testing.T) {
	db := OpenConn()
	defer db.Close()

	ctx := context.Background()
	masuk := "INSERT INTO pelanggan(id,name) VALUES(1,'Dadang')"

	_, err := db.ExecContext(ctx, masuk)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new pelanggan")
}

//DB go-databases/customer
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

//DB godb/pelanggan
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

//DB godb/user
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

//DB godb/user
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

//DB godb/user
func TestExecSQLParameter(t *testing.T) {
	db := SetConnect()
	defer db.Close()

	ctx := context.Background()

	username := "arif'; DROP TABLE user; #"
	password := "password"

	perintah := "INSERT INTO user(username, password) VALUES(?, ?)" //with VALUE (?, ?) method will reject sql INJECTION
	_, err := db.ExecContext(ctx, perintah, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new user")
}

//DB godb/comments
func TestAutoIncrementData(t *testing.T) {
	db := SetConnect()
	defer db.Close()

	ctx := context.Background()

	email := "arif@gmail.com"
	comment := "Test Comment"

	for i := 1; i <= 10; i++ {
		perintah := "INSERT INTO comments(email, comment) VALUES(?, ?)"
		result, err := db.ExecContext(ctx, perintah, email, comment)
		if err != nil {
			panic(err)
		}

		dat, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Success insert new comment with id", dat)
		fmt.Println("=======================================")
	}

}

//DB godb/comment
func TestPrepareStatement(t *testing.T) {
	db := SetConnect()
	defer db.Close()

	ctx := context.Background()
	insScript := "INSERT INTO comments(email, comment) VALUES (?, ?)"

	statement, err := db.PrepareContext(ctx, insScript)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 1; i <= 10; i++ {
		email := "arifrach" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke-" + strconv.Itoa(i) + " Test"

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment id :", id)
	}
}
