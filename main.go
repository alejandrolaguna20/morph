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

	handlers.HandlersSetup()

	portString := ":" + strconv.Itoa(s.Env.ServerPort)
	log.Fatal(http.ListenAndServe(portString, nil))
}
