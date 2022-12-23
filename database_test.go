package godatabases

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDatabases(t *testing.T) {

}

func TestOpenConn(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go-databases")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}