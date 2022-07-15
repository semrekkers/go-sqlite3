//go:build sqlite_misc
// +build sqlite_misc

package sqlite3

import (
	"github.com/semrekkers/go-sqlite3/ext/misc"
)

func init() {
	misc.LoadMiscExtensions()
}
