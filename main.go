package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/howeyc/fsnotify"
	"github.com/jessevdk/go-flags"
)

var fileIsChange int = 0
var fileIsOut int = 0

const PackageFlag string = "[Simple Go Server]"
const PackageFlagWarning string = "Warning [Simple Go Server]"

type options struct {
	Port    string `short:"p" long:"port" description:"listen port setting"`
	Content string `short:"c" long:"content" description:"content base path setting" required:"true"`
	Hot     bool   `long:"hot" description:"automatically to watch your change file"`
}

var opts options

func fileChange(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strconv.Itoa(fileIsChange)))
	fileIsChange = 0
}

func fileCallback(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
	if r.URL.Query().Get("isOK") == "true" {
		fileIsOut = 0
	}
}

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

	_, errArgs := flags.ParseArgs(&opts, os.Args)
	if errArgs != nil {
		log.Fatal(errArgs)
	}

	var port string
	//
	if len(opts.Content) != 0 && opts.Hot == true {
		watcher, errr := fsnotify.NewWatcher()
		if errr != nil {
			log.Fatal(errr)
		}

		pwd, _ := os.Getwd()
		go func() {
			for {
				select {
				case e := <-watcher.Event:
					if e.IsModify() && fileIsOut == 0 {
						fileIsChange = 1
						fileIsOut = 1
					}
				}
			}
		}()
		errs := watcher.Watch(pwd + opts.Content)
		if errs != nil {
			log.Fatal(errs)
		}

		log.Println(PackageFlag, "Watch file change on:", opts.Content)
	} else if len(opts.Content) != 0 && opts.Hot == false {
		log.Println(PackageFlagWarning + "You should use --hot start to watch file change...")
	}
	//
	if len(opts.Port) == 0 {
		port = "3000"
	} else {
		port = string(opts.Port)
	}
	http.HandleFunc("/", index)
	if opts.Hot == true {
		http.HandleFunc("/simple-go-server-file-change", fileChange)
		http.HandleFunc("/simple-go-server-file-callback", fileCallback)
	}
	log.Println(PackageFlag, "Listen on port", port+"...")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
