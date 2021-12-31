package search

import (
	"database/sql"
	"fmt"
	"html/template"
	"strings"
)

const fullTextSearchQuery = `
SELECT
workid, act, scene, description, ts_headline(body, q)
FROM (
SELECT
workid, act, scene, description, body, ts_rank(tsv, q) as rank, q
FROM
scenes, plainto_tsquery('%s') q
WHERE
tsv @@ q
ORDER BY
rank DESC
LIMIT
10
) AS Q
ORDER BY
rank DESC;`

type TextSearch struct{}

func (s *TextSearch) NewSearch(searchText string) (*SearchResults, error) {
	db, err := connectToDB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %s", err)
	}

	rows, err := s.searchDB(db, searchText)
	if err != nil {
		return nil, err
	}

	return s.formatSearchResults(rows)
}

// Search the DB using the pre-formatted SQL query and search text
func (s *TextSearch) searchDB(db *sql.DB, searchText string) (*sql.Rows, error) {
	// Query the database
	query := fmt.Sprintf(fullTextSearchQuery, searchText)
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute search query: %s", err)
	}

	return rows, nil
}

// Format the search results
func (s *TextSearch) formatSearchResults(rows *sql.Rows) (*SearchResults, error) {
	defer rows.Close()

	// Format the returned rows from the query
	var rowData []sqlRow
	for rows.Next() {
		var r sqlRow
		var snip string
		rows.Scan(&r.WorkID, &r.Act, &r.Scene, &r.Description, &snip)

		// Format the string containing HTML elements into a HTML template so it can be rendered
		r.HTMLResult = template.HTML(strings.Replace(snip, "\n", "<br>", -1))

		rowData = append(rowData, r)
	}

	return &SearchResults{rowData}, nil
}
