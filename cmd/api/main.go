package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julioc98/shion/cmd/api/handler"
	"github.com/julioc98/shion/cmd/api/router"
	"github.com/julioc98/shion/internal/app/user"
	db "github.com/julioc98/shion/internal/db"
	"github.com/julioc98/shion/pkg/env"
	"github.com/julioc98/shion/pkg/middleware"
)

func handlerHi(w http.ResponseWriter, r *http.Request) {
	msg := "Ola, Seja bem vindo ao Shion!!"
	log.Println(msg)
	w.Write([]byte(msg))
}

func main() {

	conn := db.Conn()
	defer conn.Close()
	db.Migrate(conn)

	r := mux.NewRouter()
	r.Use(middleware.Logging)

	userRep := user.NewPostgresRepository(conn)
	userService := user.NewService(userRep)
	userHandler := handler.NewUserHandler(userService)

	router.SetUserRoutes(userHandler, r.PathPrefix("/users").Subrouter())

	r.HandleFunc("/", handlerHi)
	http.Handle("/", r)

	port := env.Get("PORT", "5001")
	log.Printf(`%s listening on port: %s `, env.Get("APP", "shion"), port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
