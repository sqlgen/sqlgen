// Package models contains the types for schema 'booktest'.
package models

// GENERATED BY XOXO. DO NOT EDIT.

// AuthorBookResult is the result of a search.
type AuthorBookResult struct {
	AuthorID   int    // author_id
	AuthorName string // author_name
	BookID     int    // book_id
	BookIsbn   string // book_isbn
	BookTitle  string // book_title
	BookTags   string // book_tags
}

// AuthorBookResultsByTag runs a custom query, returning results as AuthorBookResult.
func AuthorBookResultsByTag(db XODB, tag string) ([]*AuthorBookResult, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`a.author_id AS author_id, ` +
		`a.name AS author_name, ` +
		`b.book_id AS book_id, ` +
		`b.isbn AS book_isbn, ` +
		`b.title AS book_title, ` +
		`b.tags AS book_tags ` +
		`FROM books b ` +
		`JOIN authors a ON a.author_id = b.author_id ` +
		`WHERE b.tags LIKE CONCAT('%', ?, '%')`

	// run query
	XOLog(sqlstr, tag)
	q, err := db.Query(sqlstr, tag)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AuthorBookResult{}
	for q.Next() {
		abr := AuthorBookResult{}

		// scan
		err = q.Scan(&abr.AuthorID, &abr.AuthorName, &abr.BookID, &abr.BookIsbn, &abr.BookTitle, &abr.BookTags)
		if err != nil {
			return nil, err
		}

		res = append(res, &abr)
	}

	return res, nil
}
