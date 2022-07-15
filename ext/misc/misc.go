package misc

/*
#cgo CFLAGS: -I../..

#include <sqlite3ext.h>

SQLITE_EXTENSION_INIT3

int sqlite3_series_init(
  sqlite3 *db,
  char **pzErrMsg,
  const sqlite3_api_routines *pApi
);

static int load_misc_extensions() {
	int rc = sqlite3_auto_extension((void *)sqlite3_series_init);
	return rc;
}
*/
import "C"

func LoadMiscExtensions() error {
	C.load_misc_extensions()
	return nil
}
