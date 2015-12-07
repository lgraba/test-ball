// hello.go
// A Hellow World HTTP serving example application
// Next up: Deploy a guestbook, fool

package hello

import (
	// "fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	// "appengine"
	// "appengine/user"
)

var (
	guestbookForm []byte
	signTemplate = template.Must(template.ParseFiles("guestbook.html"))
)


func init() {
	content, err := ioutil.ReadFile("guestbookform.html")
	if err != nil {
		panic(err)
	}
	guestbookForm = content

	http.HandleFunc("/", root)
	http.HandleFunc("/entry", entry)
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Write(guestbookForm)
}

func sign(w http.ResponseWriter, r *http.Request) {
	err := signTemplate.Execute(w, r.FormValue("content"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}