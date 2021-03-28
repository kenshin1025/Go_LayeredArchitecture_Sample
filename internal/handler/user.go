package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/kenshin1025/Go_LayeredArchitecture_Sample/internal/apierr"
)

//createUserのリクエストのjsonに合わせた構造体の定義
type ReqCreateUserJSON struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

//createUserのレスポンスのjsonに合わせた構造体の定義
type ResCreateUserJSON struct {
	ID int64 `json:"id"`
}

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//POSTリクエストの時だけ処理をする
		if r.Method == "POST" {
			//jsonからgoの構造体にデコードする
			var user ReqCreateUserJSON
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				log.Fatal(err)
				return
			}

			//送られてきたユーザーのEmailがすでに登録されていないかチェックする
			err := findByEmil(db, user.Email)
			//すでに登録されていたらエラーを返す
			if errors.Is(err, apierr.ErrEmailAlreadyExists) {
				log.Fatal(err)
				w.Header().Set("Content-Type", "application/json;charset=utf-8")
				w.WriteHeader(http.StatusBadRequest)
				return
			} else if err != nil {
				log.Fatal(err)
				return
			}

			//送られてきたユーザーの情報をDBに登録する
			result, err := db.Exec("INSERT INTO users(name, email) VALUES(?,?)", user.Name, user.Email)
			if err != nil {
				log.Fatal(err)
				return
			}

			// insertしたuserのidを取得
			id, err := result.LastInsertId()
			if err != nil {
				log.Fatal(err)
				return
			}

			//レスポンス用にヘッダーをセットする
			w.Header().Set("Content-Type", "application/json;charset=utf-8")
			w.WriteHeader(http.StatusCreated)

			if err := json.NewEncoder(w).Encode(&ResCreateUserJSON{
				ID: id,
			}); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func findByEmil(db *sql.DB, email string) error {
	//送られてきたユーザーのEmailがすでに登録されていないかチェックする
	err := db.QueryRow("SELECT name FROM user WHERE token = ?", email).Err()
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		return err
	}
	return apierr.ErrEmailAlreadyExists
}
