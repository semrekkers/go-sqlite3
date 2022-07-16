//go:build sqlite_misc
// +build sqlite_misc

package sqlite3

/*
#include "sqlite3-binding.h"

int sqlite3_series_init(sqlite3 *db, char **pzErrMsg, const sqlite3_api_routines *pApi);

static int load_misc_extensions() {
	return sqlite3_auto_extension((void *)sqlite3_series_init);
}
*/
import "C"

func init() {
	C.load_misc_extensions()
}
