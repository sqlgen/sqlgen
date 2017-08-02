// Package models contains the types for schema 'public'.
package models

// GENERATED BY XOXO. DO NOT EDIT.

// Proc represents a stored procedure.
type Proc struct {
	ProcName   string // proc_name
	ReturnType string // return_type
}

// PgProcs runs a custom query, returning results as Proc.
func PgProcs(db XODB, schema string) ([]*Proc, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`p.proname, ` + // ::varchar AS proc_name
		`pg_get_function_result(p.oid) ` + // ::varchar AS return_type
		`FROM pg_proc p ` +
		`JOIN ONLY pg_namespace n ON p.pronamespace = n.oid ` +
		`WHERE n.nspname = $1`

	// run query
	XOLog(sqlstr, schema)
	q, err := db.Query(sqlstr, schema)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Proc{}
	for q.Next() {
		p := Proc{}

		// scan
		err = q.Scan(&p.ProcName, &p.ReturnType)
		if err != nil {
			return nil, err
		}

		res = append(res, &p)
	}

	return res, nil
}

// MyProcs runs a custom query, returning results as Proc.
func MyProcs(db XODB, schema string) ([]*Proc, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`r.routine_name AS proc_name, ` +
		`p.dtd_identifier AS return_type ` +
		`FROM information_schema.routines r ` +
		`INNER JOIN information_schema.parameters p ` +
		`ON p.specific_schema = r.routine_schema AND p.specific_name = r.routine_name AND p.ordinal_position = 0 ` +
		`WHERE r.routine_schema = ?`

	// run query
	XOLog(sqlstr, schema)
	q, err := db.Query(sqlstr, schema)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Proc{}
	for q.Next() {
		p := Proc{}

		// scan
		err = q.Scan(&p.ProcName, &p.ReturnType)
		if err != nil {
			return nil, err
		}

		res = append(res, &p)
	}

	return res, nil
}
