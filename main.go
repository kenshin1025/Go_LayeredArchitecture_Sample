package main

import (
	"database/sql"
	"fmt"
	"log"
)

// わかりやすさのためにこうしてるけどセキュリティ的にはenvとかから取り出すようにしたほうが良いと思います
const (
	DBUserName = "root"
	DBPassword = "passw0rd"
	DBHost     = "sample_db"
	DBPort     = "3306"
	DBName     = "sample"
)

func main() {
	fmt.Printf("Starting server at 'http://localhost:8080'\n")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUserName, DBPassword, DBHost, DBPort, DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
