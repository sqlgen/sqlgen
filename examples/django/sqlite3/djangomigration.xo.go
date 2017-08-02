// Package sqlite3 contains the types for schema ''.
package sqlite3

// GENERATED BY XOXO. DO NOT EDIT.

import (
	"errors"
	"time"
)

// DjangoMigration represents a row from 'django_migrations'.
type DjangoMigration struct {
	ID      int       `json:"id"`      // id
	App     string    `json:"app"`     // app
	Name    string    `json:"name"`    // name
	Applied time.Time `json:"applied"` // applied

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the DjangoMigration exists in the database.
func (dm *DjangoMigration) Exists() bool {
	return dm._exists
}

// Deleted provides information if the DjangoMigration has been deleted from the database.
func (dm *DjangoMigration) Deleted() bool {
	return dm._deleted
}

// Insert inserts the DjangoMigration to the database.
func (dm *DjangoMigration) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if dm._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO django_migrations (` +
		`app, name, applied` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, dm.App, dm.Name, dm.Applied)
	res, err := db.Exec(sqlstr, dm.App, dm.Name, dm.Applied)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	dm.ID = int(id)
	dm._exists = true

	return nil
}

// Update updates the DjangoMigration in the database.
func (dm *DjangoMigration) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !dm._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if dm._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE django_migrations SET ` +
		`app = ?, name = ?, applied = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, dm.App, dm.Name, dm.Applied, dm.ID)
	_, err = db.Exec(sqlstr, dm.App, dm.Name, dm.Applied, dm.ID)
	return err
}

// Save saves the DjangoMigration to the database.
func (dm *DjangoMigration) Save(db XODB) error {
	if dm.Exists() {
		return dm.Update(db)
	}

	return dm.Insert(db)
}

// Delete deletes the DjangoMigration from the database.
func (dm *DjangoMigration) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !dm._exists {
		return nil
	}

	// if deleted, bail
	if dm._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM django_migrations WHERE id = ?`

	// run query
	XOLog(sqlstr, dm.ID)
	_, err = db.Exec(sqlstr, dm.ID)
	if err != nil {
		return err
	}

	// set deleted
	dm._deleted = true

	return nil
}

// DjangoMigrationByID retrieves a row from 'django_migrations' as a DjangoMigration.
//
// Generated from index 'django_migrations_id_pkey'.
func DjangoMigrationByID(db XODB, id int) (*DjangoMigration, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, app, name, applied ` +
		`FROM django_migrations ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	dm := DjangoMigration{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&dm.ID, &dm.App, &dm.Name, &dm.Applied)
	if err != nil {
		return nil, err
	}

	return &dm, nil
}
