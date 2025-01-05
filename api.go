package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	token "github.com/lukemassa/jclubtakeaways-web/internal/token"
)

var server bool

func init() {
	flag.BoolVar(&server, "server", false, "Whether or not to run the server")
	flag.Parse()
}

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":3000"
}

func HandleHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, []byte("HIHIHI"))
}

// GetToken API to get a single token
func GetToken(w http.ResponseWriter, r *http.Request) {
	tok := token.GetToken()
	res := fmt.Sprintf("{\"token\":\"%s\",\"error\":\"\"}", tok)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, res)
}
