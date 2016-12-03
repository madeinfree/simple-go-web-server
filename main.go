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
			router := strings.Split(r.URL.Path, "/")
			last := router[len(router)-1]
			ext := strings.Split(last, ".")[1]
			if ext == "js" {
				js_file, _ := ioutil.ReadFile(pwd + r.URL.Path)
				w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
				w.Write([]byte(js_file))
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
