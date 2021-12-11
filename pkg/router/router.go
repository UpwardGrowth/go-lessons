package router

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/UpwardGrowth/go-lessons/internal/sqlstore"
	"github.com/gorilla/mux"
)

type handler struct {
	AccountStore sqlstore.Account
}

func Activate(router *mux.Router, conn *sql.DB) {
	accountStore := sqlstore.New(conn)
	newHttpServer(router, *accountStore)
}

func newHttpServer(router *mux.Router, ac sqlstore.Account) {
	hs := handler{
		AccountStore: ac,
	}

	router.HandleFunc("/getuser/{id}", hs.getUserHandle).Methods("GET")
}

func (hs *handler) getUserHandle(w http.ResponseWriter, req *http.Request) {
	id, ok := mux.Vars(req)["id"]
	if !ok {
		http.Error(w, "no id found in request", http.StatusBadRequest)
		return
	}

	user, err := hs.AccountStore.QueryUserinfo(id)
	if err == sql.ErrNoRows {
		fmt.Println("getUserHandle QueryUserinfo err: no data")

		resp := map[string]interface{}{
			"msg": "ID not found",
		}
		response(w, resp)
		return
	}
	if err != nil {
		fmt.Printf("getUserHandle QueryUserinfo err:%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]interface{}{
		"name":  user.Name,
		"email": user.Email,
		"msg":   "ok",
	}
	fmt.Printf("getUserHandle: id %s, name %s, email %s", id, user.Name, user.Email)
	response(w, resp)
}

func response(w http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
