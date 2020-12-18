package gofocused

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type studentSQL struct {
	ID    string
	Name  string
	Age   int
	Grade int
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:sastra@tcp(172.19.0.2:3306)/gofocused")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func SQLQuery() {
	db, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 27
	rows, err := db.Query("SELECT id, name, grade FROM tb_student WHERE age = ?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []studentSQL

	for rows.Next() {
		var each = studentSQL{}
		var err = rows.Scan(&each.ID, &each.Name, &each.Grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.Name)
	}
}
