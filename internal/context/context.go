package context

import (
	"database/sql"
	"fmt"
)

type Context struct {
	db *sql.DB
}

var context Context

func (c *Context) GetDB() (db *sql.DB, e error) {
	if c.db != nil {
		db = c.db
	} else {
		e = fmt.Errorf("no sql client")
	}

	return
}
