package handler

import (
	"encoding/json"
	"go/pkg/domain"
	"go/pkg/usecase/task"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func GetAllTask(service task.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tasks []*domain.Task
		var err error
		//サービス層を呼びリストを取得
		tasks, err = service.GetAllTask()
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error reading tasks"))
			return
		}

		if tasks == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Error reading tasks"))
			return
		}

		//JSON形式変換
		var toJ []*domain.Jtask
		for _, t := range tasks {
			toJ = append(toJ, &domain.Jtask{
				ID:   t.ID,
				Name: t.Name,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error reading tasks"))
		}
	})
}

func GetTask(service task.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//パラメータ取得
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error reading tasks"))
			return
		}
		//サービス層呼びだし
		task, err := service.GetTask(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error reading tasks"))
			return
		}

		if err == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Error reading tasks"))
			return
		}
		toJ := &domain.Jtask{
			ID:   task.ID,
			Name: task.Name,
		}

		//JSON変換
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error reading tasks"))
		}
	})
}

func createTask(service task.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errMessage := "Error adding task"
		var input struct {
			ID   int `json:"ID,string"`
			Name string
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errMessage))
			return
		}
		id, err := service.CreateTask(input.ID, input.Name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errMessage))
			return
		}
		toJ := &domain.Jtask{
			ID:   id,
			Name: input.Name,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errMessage))
			return
		}
	})
}

func updateTask(service task.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errMessage := "Error reading tasks"
		//パラメータ取得
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errMessage))
			return
		}
		//レコード抽出
		/*
			task, err := service.GetTask(id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errMessage))
				return
			}
			if err == nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(errMessage))
				return
			}
		*/
		//更新
		var input domain.Task
		err2 := json.NewDecoder(r.Body).Decode(&input)
		if err2 != nil {
			log.Println(err2.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errMessage))
			return
		}
		input.ID = id
		//_ = service.UpdateTask(input)
	})
}

func TaskHandlers(r *mux.Router, n negroni.Negroni, service task.UseCase) {
	r.Handle("/task", n.With(
		negroni.Wrap(GetAllTask(service)),
	)).Methods("GET", "OPTIONS").Name("GetAllTask")

	r.Handle("/task/{id}", n.With(
		negroni.Wrap(GetTask(service)),
	)).Methods("GET", "OPTIONS").Name("GetTask")

	r.Handle("/task", n.With(
		negroni.Wrap(createTask(service)),
	)).Methods("POST", "OPTIONS").Name("createTask")

	r.Handle("/task/{id}", n.With(
		negroni.Wrap(updateTask(service)),
	)).Methods("POST", "OPTIONS").Name("updateTask")
}
