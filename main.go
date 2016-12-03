package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	pwd, _ := os.Getwd()
	switch r.Method {
	case "GET":
		if r.URL.Path == "" || r.URL.Path == "/" {
			index_template, _ := ioutil.ReadFile(pwd + "/index.html")
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write([]byte(index_template))
		}
		if r.URL.Path != "/" {
			var ext string
			router := strings.Split(r.URL.Path, "/")
			last := router[len(router)-1]
			ext_split := strings.Split(last, ".")
			if len(ext_split) > 1 {
				ext = strings.Split(last, ".")[1]
			} else {
				ext = ""
			}
			if ext == "js" {
				js_file, err := ioutil.ReadFile(pwd + r.URL.Path)
				if err != nil {
					http.NotFound(w, r)
				} else {
					w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
					w.Write([]byte(js_file))
				}
			} else {
				http.NotFound(w, r)
			}
		}
	}
}

func main() {
	var port string
	if len(os.Args) <= 1 {
		port = "3000"
	} else {
		port = string(os.Args[1])
	}
	http.HandleFunc("/", index)
	fmt.Printf("[Simple Go Server] Listen on port %s ...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
