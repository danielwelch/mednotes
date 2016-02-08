package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	// "path/filepath"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", IndexHandler)

	log.Println("Registering the following directory handlers...")
	// register a handler (FileServer) for each directory
	dirs, _ := DirNames(true)
	for _, dir := range dirs {
		log.Println(dir)
		r.PathPrefix(dir).Handler(
			http.StripPrefix(
				dir,
				http.FileServer(http.Dir("."+dir)),
			),
		)
	}

	return r
}

// DirNames returns the name of all dirs in the root directory.
// DirNames assumes that all dirs in the root directory should be the unmodified result
// of exporting a Quiver Notebook into the root directory.
// If pre == true, name of dir will be preceeded with "/"
func DirNames(pre bool) ([]string, error) {
	var dirs []string
	files, err := ioutil.ReadDir("./")
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if f.IsDir() {
			if pre {
				dirs = append(dirs, "/"+f.Name()+"/")
			} else {
				dirs = append(dirs, f.Name()+"/")
			}

		}
	}
	return dirs, nil
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	dirs, err := DirNames(false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, dirs)
}
