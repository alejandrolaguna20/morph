package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/alejandrolaguna20/morph/handlers"
	"github.com/alejandrolaguna20/morph/state"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	s := state.LoadState()

	q, err := s.Database.Query("SELECT * FROM urls")
	defer q.Close()

	if err != nil {
		log.Fatal("[ERROR] A connection could not be made")
	}

	handlers.HandlersSetup(&s)

	portString := ":" + strconv.Itoa(s.Env.ServerPort)
	log.Println("[INFO] SERVER IS UP")
	log.Fatal(http.ListenAndServe(portString, nil))
}
