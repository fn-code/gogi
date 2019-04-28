package main

import (
	"fmt"
	"log"

	"github.com/fn-code/gogi/internal/view"
	"github.com/fn-code/gogi/internal/web/server"
)

func main() {
	view.InitTemplate("./templates", ".gohtml")

	fmt.Println("Server is running")
	srv, err := server.New(":8888", "public")
	if err != nil {
		log.Fatal(err)
	}

	err = srv.Run()
	if err != nil {
		log.Fatal(err)
	}
	defer srv.Close()
}
