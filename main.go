package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kenshin1025/Go_LayeredArchitecture_Sample/internal/handler"
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
	//apiの起動確認
	fmt.Printf("Starting server at 'http://localhost:8080'\n")

	//sql.OpenをするためのdataSourceNameの生成
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUserName, DBPassword, DBHost, DBPort, DBName)

	//生成したdsnを元にsql.Open
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", test)
	http.HandleFunc("/user/create", handler.CreateUser(db))
	http.ListenAndServe(":8080", nil)
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK!")
}
