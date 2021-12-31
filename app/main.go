package main

import (
	"CurtisM132/FullTextSearchDemo/search"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	t := template.Must(template.ParseGlob("html/*"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "index", nil)
	})

	// Host CSS files so the HTML templates can access them
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))

	http.HandleFunc("/search", handleSearch)

	http.ListenAndServe(":8080", nil)
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var textSearch search.TextSearch
	sr, err := textSearch.NewSearch(r.Form["search"][0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	results := *sr

	t := template.Must(template.ParseFiles("html/search-results.html"))
	t.ExecuteTemplate(w, "searchResults", results)
}
