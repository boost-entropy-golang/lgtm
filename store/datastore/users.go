package datastore

import (
	"github.com/go-gitea/lgtm/model"

	"github.com/russross/meddler"
)

func (db *datastore) GetUser(id int64) (*model.User, error) {
	var usr = new(model.User)
	var err = meddler.Load(db, userTable, usr, id)
	return usr, err
}

func (db *datastore) GetUserLogin(login string) (*model.User, error) {
	var usr = new(model.User)
	var err = meddler.QueryRow(db, usr, rebind(userLoginQuery), login)
	return usr, err
}

func (db *datastore) CreateUser(user *model.User) error {
	return meddler.Insert(db, userTable, user)
}

func (db *datastore) UpdateUser(user *model.User) error {
	return meddler.Update(db, userTable, user)
}

func (db *datastore) DeleteUser(user *model.User) error {
	var _, err = db.Exec(rebind(userDeleteStmt), user.ID)
	return err
}

const userTable = "users"

const userLoginQuery = `
SELECT *
FROM users
WHERE user_login = ?
LIMIT 1
`

const userDeleteStmt = `
DELETE FROM users
WHERE user_id = ?
`
