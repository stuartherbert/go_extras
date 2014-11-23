// Copyright (c) 2014-present Stuart Herbert
// Released under a 3-clause BSD license
package extrafmt

import (
	"bytes"
	"fmt"
)

// Sprintnln() emulates fmt.Sprintln()'s append behaviour, something that
// fmt.Sprint(args...) does not give us :(
func Sprintnln(args ...interface{}) string {
	if len(args) == 1 {
		return fmt.Sprint(args...)
	} else {
		retval := new(bytes.Buffer)
		appendSpace := false
		for _, arg := range args {
			if appendSpace {
				retval.WriteByte(' ')
				retval.WriteString(fmt.Sprint(arg))
			} else {
				retval.WriteString(fmt.Sprint(arg))
				appendSpace = true
			}
		}

		return retval.String()
	}
}
