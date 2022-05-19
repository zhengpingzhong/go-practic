package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

func main() {
	db, err := sql.Open("mysql", "root:Root&20200306@(mysql:33306)/eventdispatch?parseTime=true")
	fmt.Println(err)
	err = db.Ping()
	fmt.Println(err)

	// create table
	//query := `
	//CREATE TABLE test (
	//    id INT AUTO_INCREMENT,
	//    username TEXT NOT NULL,
	//    password TEXT NOT NULL,
	//    created_at DATETIME,
	//    PRIMARY KEY (id)
	//);`
	//
	//_, err = db.Exec(query)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//插入数据
	username := "alex"
	password := "alex"
	createdAt := time.Now()
	result, err := db.Exec(`insert into test(username,password,created_at) values (?,?,?)`, username, password, createdAt)
	userID, err := result.LastInsertId()
	fmt.Println(userID)

	{ // Query a single user
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)

		query := "SELECT id, username, password, created_at FROM test WHERE id = ?"
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, username, password, createdAt)
	}

	{ // Query all users
		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		rows, err := db.Query(`SELECT id, username, password, created_at FROM test`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []user
		for rows.Next() {
			var u user

			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v", users)
	}

	{
		_, err := db.Exec(`DELETE FROM test WHERE id = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}
