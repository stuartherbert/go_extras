// Copyright (c) 2014-present Stuart Herbert
// Released under a 3-clause BSD license
package xreflect

import (
	"bytes"
	"reflect"
)

func get_type(val interface{}) string {
	// this will hold our return value
	buf := bytes.Buffer{}

	// what do we have?
	t := reflect.TypeOf(val)

	pkgName := t.PkgPath()
	if len(pkgName) > 0 {
		buf.WriteString(pkgName)
		buf.WriteString(".")
	}

	buf.WriteString(t.Name())

	// all done
	return buf.String()
}
