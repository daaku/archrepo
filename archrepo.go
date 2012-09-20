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
	addr = flag.String("addr", ":2724", "server address")
)

func main() {
	flag.Parse()
	err := gracehttp.Serve(
		gracehttp.Handler{*addr, http.FileServer(http.Dir(*repo))},
	)
	if err != nil {
		log.Fatal(err)
	}
}
