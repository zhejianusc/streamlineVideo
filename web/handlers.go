package web

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)
type HomePage struct {
	Name string
}
func homeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	p := &HomePage{Name :"zhejian"}
	t,e := template.ParseFiles("./template/home.html")
	if e != nil {
		log.Print("Parsing template error : %s", e)
	}
	t.Execute(w, p)
	return
}
