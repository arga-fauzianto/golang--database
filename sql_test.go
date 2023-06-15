package godatabase

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"
)

// func TestExecSql(t *testing.T) {
// 	db := GetConnection()

// 	defer db.Close()

// 	ctx := context.Background()

// 	script := "SELECT id, name FROM customer"

// 	rows, err := db.QueryContext(ctx, script)

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var id, name string
// 		err = rows.Scan(&id, &name)

// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println("Id:", id)
// 		fmt.Println("Name:", name)
// 	}

// }

func TestQueryComplex(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	syntax := "SELECT id, name, email, balance, rating, birth_date, created_at , married FROM customer"

	rows, err := db.QueryContext(ctx, syntax)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name, email string
		var balance int32
		var rating float64
		var birthDate, createdAt time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &createdAt, &married)

		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id, "Name:", name, "Email:", email, "Balance:", balance,
			"Rating:", rating, "Birth_Date:", birthDate, "Married:", married,
			"CreatedAt:", createdAt)
	}

}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "arga@gmail.com"
	comment := "ini comment satu"

	script := "INSERT INTO comments(email, comment) VALUES(? , ?)"
	result, err := db.ExecContext(ctx, script, email, comment)

	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	fmt.Println("comment by id", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO comments(email, comment) VALUES(? , ?)"
	statement, err := db.PrepareContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 0; i > 10; i++ {

		email := "arga" + strconv.Itoa(i) + "@gmail.com"
		comment := "komentar ke " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)

		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("comment id", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments(email, comment) VALUES(? , ?)"

	for i := 0; i > 10; i++ {

		email := "arga" + strconv.Itoa(i) + "@gmail.com"
		comment := "komentar ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)

		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("comment id", id)
	}

	err = tx.Commit()

	if err != nil {
		panic(err)
	}

}
