package main

import (
	"CurtisM132/FullTextSearchDemo/search"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	// Serve index.html file
	t := template.Must(template.ParseGlob("html/*"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "index", nil)
	})

	// Host CSS files so the HTML templates can access them
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))

	http.HandleFunc("/search", handleSearch)

	http.ListenAndServe(":8080", nil)
}

// Handles a POST request, creates a new search, and displays the search results in a HTML template
func handleSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()

		var textSearch search.TextSearch
		sr, err := textSearch.NewSearch(r.Form["search"][0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		t := template.Must(template.ParseFiles("html/search-results.html"))
		t.ExecuteTemplate(w, "searchResults", *sr)
	} else {
		http.Error(w, "endpoint only accepts POSTS", http.StatusInternalServerError)
	}
}
