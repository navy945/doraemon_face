package dammy

import (
	"database/sql"
	"encoding/json"

	//"go/pkg/interface/handler"
	"log"
	"net/http"
	//"github.com/gorilla/mux"
	//"github.com/urfave/negroni"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	var dbPool *sql.DB
	// Dbの初期化
	dbPool, err := sql.Open("pgx", "host=localhost user=postgres password=postgres dbname=tododev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer dbPool.Close()

	//パラメータ取得
	/*
		vars := mux.Vars(r)
		//queryID, _ := strconv.Atoi(vars["id"])
		queryID, _ := vars["id"]
		log.Println(queryID)
	*/

	sql := "SELECT id, email FROM t_user WHERE id=$1;"
	stmt, err := dbPool.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	// 検索結果格納用の TestUser
	type User struct {
		ID    int    `json:"id"`
		Email string `json:"email"`
	}
	var user User

	queryID := 1
	// queryID を埋め込み SQL の実行、検索結果1件の取得
	err = stmt.QueryRow(queryID).Scan(&user.ID, &user.Email)
	if err != nil {
		log.Fatal(err)
	}
	res := &User{ID: user.ID, Email: user.Email}

	json, _ := json.Marshal(res)
	log.Println(res)

	w.Write(json)
}
