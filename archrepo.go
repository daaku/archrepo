// Command archrepo provides a socket activatable http server for
// serving a arch package repository.
package main

import (
	"flag"
	"github.com/daaku/go.grace/gracehttp"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	repo = flag.String(
		"repo",
		filepath.Join(os.Getenv("HOME"), "pkgs", "repo"),
		"repository directory")
	repoAddr  = flag.String("addr", ":2724", "repo server address")
	adminAddr = flag.String("admin", ":2725", "admin server address")
)

func rebuild(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world\n"))
}

func adminHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/rebuild", rebuild)
	return mux
}

func main() {
	flag.Parse()
	err := gracehttp.Serve(
		gracehttp.Handler{*repoAddr, http.FileServer(http.Dir(*repo))},
		gracehttp.Handler{*adminAddr, adminHandler()},
	)
	if err != nil {
		log.Fatal(err)
	}
}
