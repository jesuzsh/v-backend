package main

import (
	"database"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

// templates calls ParseFiles on intialization to prevent calling the function
// every time a page is rendered. ExecuteTemplate method can now be used to
// render any of the specified templates.
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

// validPath is a global variable that stores our validation expression. It
// prevents the user from supplying an arbitrary path to be read/written on the
// server.
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

type Page struct {
	Title string
	Body  []byte
}

// Transaction represents the basic unit of what is being reported.
type Transaction struct {
	ID           int
	Name         string
	AmountUSD    int
	CardName     string
	CategoryName string
}

// Card represents what's used for payments.
type Card struct {
	ID    int
	Name  string
	Owner string
}

// Category is the classifcation that every transaction should fall under.
type Category struct {
	ID          int
	Name        string
	IsEssential bool
}

// save is a method of each struct. The job is to persist the data in the
// database. Is this an interface.
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// loadPage gets the page content that is to be rendered for the user. The
// content is stored as a struct and is returned from this function.
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// renderTemplate ensures that a template is rendered for the user. At the
// moment renderTemplate has http.ResponseWriter as a parameter.
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// viewHandler will show the user the exact view they'd like to see given their
// most recent action. This function needs to handle the traffic going between
// anything working
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	res := database.Save(3)

	fmt.Print(res)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
