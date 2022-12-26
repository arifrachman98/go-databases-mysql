package repository

import (
	"context"
	"fmt"
	"testing"

	
	"github.com/arifrachman98/go-databases-mysql/entity"
	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	comentRepo := NewCommentRepo(gdb.SetConnect())

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
