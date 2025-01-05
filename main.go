package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lukemassa/jclubtakeaways-web/internal/templates"
	"github.com/lukemassa/jclubtakeaways-web/internal/token"
)

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":3000"
}

func usage() {
	log.Fatal("Usage: server|templates|token")
}

func main() {

	templater := templates.NewFromDirectory("templates")
	tokener := token.New()
	if len(os.Args) != 2 {
		usage()
	}
	switch os.Args[1] {
	case "server":
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/index.html", http.StatusMovedPermanently)
		})
		http.Handle("/{page}", templater)
		http.Handle("/token", tokener)
		port := getPort()
		fmt.Printf("Listening on %s", port)

		http.ListenAndServe(port, nil)
	case "templates":
		err := templater.Write("docs")
		if err != nil {
			log.Fatal(err)
		}
	case "token":
		fmt.Print(tokener.Get())
		fmt.Println()
	default:
		usage()
	}

}
