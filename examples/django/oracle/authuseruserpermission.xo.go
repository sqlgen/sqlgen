// Package oracle contains the types for schema 'django'.
package oracle

// GENERATED BY XOXO. DO NOT EDIT.

import "errors"

// AuthUserUserPermission represents a row from 'django.auth_user_user_permissions'.
type AuthUserUserPermission struct {
	ID           float64 // id
	UserID       float64 // user_id
	PermissionID float64 // permission_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the AuthUserUserPermission exists in the database.
func (auup *AuthUserUserPermission) Exists() bool {
	return auup._exists
}

// Deleted provides information if the AuthUserUserPermission has been deleted from the database.
func (auup *AuthUserUserPermission) Deleted() bool {
	return auup._deleted
}

// Insert inserts the AuthUserUserPermission to the database.
func (auup *AuthUserUserPermission) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if auup._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO django.auth_user_user_permissions (` +
		`user_id, permission_id` +
		`) VALUES (` +
		`:1, :2` +
		`) RETURNING id /*lastInsertId*/ INTO :pk`

	// run query
	XOLog(sqlstr, auup.UserID, auup.PermissionID, nil)
	res, err := db.Exec(sqlstr, auup.UserID, auup.PermissionID, nil)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	auup.ID = float64(id)
	auup._exists = true

	return nil
}

// Update updates the AuthUserUserPermission in the database.
func (auup *AuthUserUserPermission) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !auup._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if auup._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE django.auth_user_user_permissions SET ` +
		`user_id = :1, permission_id = :2` +
		` WHERE id = :3`

	// run query
	XOLog(sqlstr, auup.UserID, auup.PermissionID, auup.ID)
	_, err = db.Exec(sqlstr, auup.UserID, auup.PermissionID, auup.ID)
	return err
}

// Save saves the AuthUserUserPermission to the database.
func (auup *AuthUserUserPermission) Save(db XODB) error {
	if auup.Exists() {
		return auup.Update(db)
	}

	return auup.Insert(db)
}

// Delete deletes the AuthUserUserPermission from the database.
func (auup *AuthUserUserPermission) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !auup._exists {
		return nil
	}

	// if deleted, bail
	if auup._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM django.auth_user_user_permissions WHERE id = :1`

	// run query
	XOLog(sqlstr, auup.ID)
	_, err = db.Exec(sqlstr, auup.ID)
	if err != nil {
		return err
	}

	// set deleted
	auup._deleted = true

	return nil
}

// AuthPermission returns the AuthPermission associated with the AuthUserUserPermission's PermissionID (permission_id).
//
// Generated from foreign key 'd86aa2ae91c20ae0c7ed11fa07e6da'.
func (auup *AuthUserUserPermission) AuthPermission(db XODB) (*AuthPermission, error) {
	return AuthPermissionByID(db, auup.PermissionID)
}

// AuthUser returns the AuthUser associated with the AuthUserUserPermission's UserID (user_id).
//
// Generated from foreign key 'd9da0966de759042d839a70a2bd885'.
func (auup *AuthUserUserPermission) AuthUser(db XODB) (*AuthUser, error) {
	return AuthUserByID(db, auup.UserID)
}

// AuthUserUserPermissionByUserIDPermissionID retrieves a row from 'django.auth_user_user_permissions' as a AuthUserUserPermission.
//
// Generated from index 'auth_use_user_id_14a6b632_uniq'.
func AuthUserUserPermissionByUserIDPermissionID(db XODB, userID float64, permissionID float64) (*AuthUserUserPermission, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, user_id, permission_id ` +
		`FROM django.auth_user_user_permissions ` +
		`WHERE user_id = :1 AND permission_id = :2`

	// run query
	XOLog(sqlstr, userID, permissionID)
	auup := AuthUserUserPermission{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, userID, permissionID).Scan(&auup.ID, &auup.UserID, &auup.PermissionID)
	if err != nil {
		return nil, err
	}

	return &auup, nil
}

// AuthUserUserPermissionsByUserID retrieves a row from 'django.auth_user_user_permissions' as a AuthUserUserPermission.
//
// Generated from index 'auth_user_user_permissions1cca'.
func AuthUserUserPermissionsByUserID(db XODB, userID float64) ([]*AuthUserUserPermission, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, user_id, permission_id ` +
		`FROM django.auth_user_user_permissions ` +
		`WHERE user_id = :1`

	// run query
	XOLog(sqlstr, userID)
	q, err := db.Query(sqlstr, userID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AuthUserUserPermission{}
	for q.Next() {
		auup := AuthUserUserPermission{
			_exists: true,
		}

		// scan
		err = q.Scan(&auup.ID, &auup.UserID, &auup.PermissionID)
		if err != nil {
			return nil, err
		}

		res = append(res, &auup)
	}

	return res, nil
}

// AuthUserUserPermissionsByPermissionID retrieves a row from 'django.auth_user_user_permissions' as a AuthUserUserPermission.
//
// Generated from index 'auth_user_user_permissions8afe'.
func AuthUserUserPermissionsByPermissionID(db XODB, permissionID float64) ([]*AuthUserUserPermission, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, user_id, permission_id ` +
		`FROM django.auth_user_user_permissions ` +
		`WHERE permission_id = :1`

	// run query
	XOLog(sqlstr, permissionID)
	q, err := db.Query(sqlstr, permissionID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AuthUserUserPermission{}
	for q.Next() {
		auup := AuthUserUserPermission{
			_exists: true,
		}

		// scan
		err = q.Scan(&auup.ID, &auup.UserID, &auup.PermissionID)
		if err != nil {
			return nil, err
		}

		res = append(res, &auup)
	}

	return res, nil
}

// AuthUserUserPermissionByID retrieves a row from 'django.auth_user_user_permissions' as a AuthUserUserPermission.
//
// Generated from index 'sys_c004985'.
func AuthUserUserPermissionByID(db XODB, id float64) (*AuthUserUserPermission, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, user_id, permission_id ` +
		`FROM django.auth_user_user_permissions ` +
		`WHERE id = :1`

	// run query
	XOLog(sqlstr, id)
	auup := AuthUserUserPermission{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&auup.ID, &auup.UserID, &auup.PermissionID)
	if err != nil {
		return nil, err
	}

	return &auup, nil
}
