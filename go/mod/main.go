package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const INDEX = `<!DOCTYPE html>
<html>
  <head>
    <title>Powered By Paketo Buildpacks</title>
  </head>
  <body style="bgcolor:green">
  <br>asfhafih<br>
    <img style="display: block; margin-left: auto; margin-right: auto; width: 50%;" src="https://assets.uni-wuppertal.de/dev/relaunch-assets/Resources//Public/Images/logo_header_white.svg"></img>
  </body>
</html>`

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, INDEX)
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
