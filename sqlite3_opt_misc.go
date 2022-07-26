//go:build sqlite_misc
// +build sqlite_misc

package sqlite3

/*
#define SQLITE_CORE

#include "sqlite3-binding.h"
#include "sqlite3ext.h"

#include "./ext/series.c"

static int load_misc_extensions() {
	return sqlite3_auto_extension((void *)sqlite3_series_init);
}
*/
import "C"

func init() {
	C.load_misc_extensions()
}
