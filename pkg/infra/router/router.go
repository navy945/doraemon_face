package router

//db接続
//ハンドラへのルーティング

import (
	"database/sql"
	"go/pkg/interface/handler"
	"go/pkg/usecase/task"
	"log"
	"net/http"

	"go/pkg/infra/repository"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/jackc/pgx/v5"
)

func InitRouter() {
	var dbPool *sql.DB
	// Dbの初期化
	dbPool, err := sql.Open("pgx", "host=localhost user=postgres password=postgres dbname=tododev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer dbPool.Close()

	taskRepo := repository.NewDbAccess(dbPool)
	taskService := task.NewService(taskRepo)

	//ルーター初期化
	router := mux.NewRouter()

	n := negroni.New(
		//negroni.HandlerFunc(cors.Cors),
		negroni.NewLogger(),
	)

	handler.TaskHandlers(router, *n, taskService)
	http.Handle("/", router)

	/*
		server := http.Server{
			Addr: "127.0.0.1:8080",
		}
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	*/
	srv := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: context.ClearHandler(http.DefaultServeMux),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err.Error())
	}
}
