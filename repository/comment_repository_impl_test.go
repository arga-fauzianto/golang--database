package repository

import (
	"context"
	"fmt"
	learn_go_database "learn_goDatabase"
	"learn_goDatabase/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {

	commentRepository := NewCommentRepository(learn_go_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "arga@test.com",
		Comment: "ini Comment satu",
	}

	result, err := commentRepository.Insert(ctx, comment)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
