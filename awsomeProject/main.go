package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Person struct {
	Name string
	Age  int
}

var Dbconnection *sql.DB

func main() {
	Dbconnection, _ := sql.Open("sqlite3", "E:/Go/example.sql")

	defer Dbconnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS person(
		name STRING,
		age INT)`
	_, err := Dbconnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	// _, err = Dbconnection.Exec(cmd, "Nancy", 24)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cmd = "UPDATE person set age = ? where name = ?"
	// Dbconnection.Exec(cmd, 35, "Mike")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cmd = "select * from person"
	// rows, _ := Dbconnection.Query(cmd)
	// defer rows.Close()
	// var pp []Person
	// for rows.Next() {
	// 	var p Person
	// 	err := rows.Scan(&p.Name, &p.Age)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	pp = append(pp, p)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// for _, p := range pp {
	// 	fmt.Println(p.Name, p.Age)
	// }

	// cmd = "select * from person where age = ?"
	// row := Dbconnection.QueryRow(cmd, 24)
	// var p Person
	// err = row.Scan(&p.Name, &p.Age)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Println("No row")
	// 	} else {
	// 		log.Println(err)
	// 	}
	// }
	// fmt.Println(p.Name, p.Age)

	// cmd = "delete from person where name = ?"
	// _, err = Dbconnection.Exec(cmd, "Nancy")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	tableName := "person"
	cmd = fmt.Sprintf("select * from %s", tableName)
	rows, _ := Dbconnection.Query(cmd)
	defer rows.Close()
	var pp []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}
}
