// +build sqlite3_misc

package sqlite3

/*
#define SQLITE_CORE

#ifndef USE_LIBSQLITE3
#include "sqlite3-binding.h"
#else
#include <sqlite3.h>
#endif

#include "./misc/series.c"

static int load_misc_extensions() {
	return sqlite3_auto_extension((void *)sqlite3_series_init);
}
*/
import "C"

func init() {
	C.load_misc_extensions()
}
