package search

import "html/template"

type SearchResults struct {
	Rows []sqlRow
}

type sqlRow struct {
	WorkID      string
	Act         int
	Scene       int
	Description string
	HTMLResult  template.HTML
}
