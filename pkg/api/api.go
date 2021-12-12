package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/UpwardGrowth/go-lessons/pkg/db"
	httpServer "github.com/UpwardGrowth/go-lessons/pkg/router"
	"github.com/gorilla/mux"
)

var (
	dbHost     = "127.0.0.1"
	dbPort     = 3306
	dbUser     = "root"
	dbPassword = "123456"
	dbName     = "demo"

	serverAddr = ":8080"
)

func Start() {
	// init db
	conn, err := db.GetConnction(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		fmt.Printf("api: db stack trace:\n%+v\v", err)
		panic(err)
	}

	defer conn.Close()

	// init http server
	router := mux.NewRouter()
	httpServer.Activate(router, conn)

	srv := &http.Server{
		Addr:    serverAddr,
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Println("api: ListenAndServe err\n", err)
		panic(err)
	}
}
