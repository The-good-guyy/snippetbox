package main
import (
"errors"
"fmt"
"net/http"
"strconv"
"html/template"
"snippetbox.hientt/internal/models"
)
func (app *application)home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	snippets,err:= app.snippets.Latest()
	if err!=nil{
		app.serverError(w,err)
		return
	}
	files:=[]string{"./ui/html/base.tmpl.html","./ui/html/pages/home.tmpl.html","./ui/html/partials/nav.tmpl.html"}

	ts,err:= template.ParseFiles(files...)
	if err != nil{
		app.serverError(w,err)
		return
	}
	data:=&templateData{Snippets:snippets,}
	err = ts.ExecuteTemplate(w,"base",data)
	if err!= nil{
		app.serverError(w,err)
	}
}
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	snippet,err:=app.snippets.Get(id)
	if err!=nil{
		if errors.Is(err,models.ErrNoRecord){
			app.notFound(w)
		} else{
			app.serverError(w,err)
		}
		return
	}
	files:=[]string{"./ui/html/base.tmpl.html","./ui/html/pages/view.tmpl.html","./ui/html/partials/nav.tmpl.html"}
	ts,err := template.ParseFiles(files...)
	if err!=nil{
		app.serverError(w,err)
		return
	}
	data:=&templateData{Snippet:snippet}
	err = ts.ExecuteTemplate(w,"base",data)
	if err!=nil{
		app.serverError(w,err)
	}
}
func (app *application)snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	title:= "0 snail"
	content:= "0 snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires:=7


	id,err:= app.snippets.Insert(title,content,expires)
	if err!=nil{
		app.serverError(w,err)
		return
	}
	http.Redirect(w,r,fmt.Sprintf("/snippet?id=%d",id),http.StatusSeeOther)
}