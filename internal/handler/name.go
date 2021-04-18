package handler

import (
	"awesomeProject1/internal/handler/dto"
	"awesomeProject1/internal/service"

	"encoding/json"
	"log"
	"net/http"
)

type Name interface {
	GetName(nameManager service.NameService) func(w http.ResponseWriter, r *http.Request)
	SetName(nameManager service.NameService) func(w http.ResponseWriter, r *http.Request)
	RemoveName(nameManager service.NameService) func(w http.ResponseWriter, r *http.Request)
}

type NameHandler struct {
	NameManager service.NameService
}

func InitName(nameService service.NameService) {
	n := NameHandler{
		NameManager: nameService,
	}
	// я не знаю, почему роуты настолько странные
	http.HandleFunc("/get/name", n.GetName())
	http.HandleFunc("/set/name", n.SetName())
	http.HandleFunc("/remove/name", n.RemoveName())
}

func (nh *NameHandler) GetName() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		q := r.URL.Query()
		id := q.Get("id")
		if id == "" {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		userName, err := nh.NameManager.Get(id)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		n := dto.NameResponse{
			Name: userName,
		}

		b, err := json.Marshal(n)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func (nh *NameHandler) SetName() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		id := q.Get("id")
		newName := q.Get("name")
		if id == "" || newName == "" {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		user := dto.SetNameRequest{
			ID:   id,
			Name: newName,
		}
		err := nh.NameManager.Set(user)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("done"))
	}
}

func (nh *NameHandler) RemoveName() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		id := q.Get("id")
		if id == "" {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		err := nh.NameManager.Remove(id)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("done"))
	}
}
