package repository

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/arifrachman98/go-databases-mysql/entity"
	_ "github.com/go-sql-driver/mysql"
)

func open() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/godb")
	if err != nil {
		panic(err)
	}

	return db
}

func TestCommentInsert(t *testing.T) {

	comentRepo := NewCommentRepo(open())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "TestedRepo@gmail.com",
		Comment: "Test Comment Repository",
	}

	result, err := comentRepo.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
