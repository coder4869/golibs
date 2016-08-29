// Copyright (c) 2016 coder4869 ( https://github.com/coder4869/golibs ). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*	glio/glprint.go
 */

package glio

import (
	"fmt"
	"io"

	"github.com/coder4869/golibs/gllog"
)

const (
	PRINTF_FMT     int = 0x1                     //fmt print to terminal
	PRINTF_LOG     int = 0x4                     //log print to log file
	FPRINTF_FMT    int = 0x8                     //fmt print to client
	PRINTF_FMT_LOG int = PRINTF_FMT | PRINTF_LOG //fmt print to terminal and log print to log file
)

/*
flag is print option, including PRINTF_FMT, PRINTF_LOG, FPRINTF_FMT
when FPRINTF_FMT option is used, the parameter writer can't be nil;
when FPRINTF_FMT option is not used, the parameter writer should be nil;
*/
func Printf(flag int, w io.Writer, format string, a ...interface{}) {

	io.MultiWriter()

	if flag&PRINTF_FMT != 0 { //fmt printf to terminal
		fmt.Printf(format, a...)
	}

	if flag&PRINTF_LOG != 0 { //log printf to log file
		//		gllog.SetCurrentFileName()
		gllog.AppendLog(fmt.Sprintf(format, a...))
	}

	if flag&FPRINTF_FMT != 0 { //fmt printf to client
		if w == nil {
			gllog.AppendLog("parameter writer error")
			return
		}
		fmt.Fprintf(w, format, a...)
	}
}

/*
fmt print to terminal and log print to log file
which means:
 flag = PRINTF_FMT | PRINTF_LOG (or PRINTF_FMT_LOG)
 w = nil
*/
func FLPrintf(format string, a ...interface{}) {
	Printf(PRINTF_FMT|PRINTF_LOG, nil, format, a...)
}

/*
fmt print to terminal and log print to log file and fmt print to client
which means:
 flag = PRINTF_FMT | PRINTF_LOG | FPRINTF_FMT
*/
func FFLPrintf(w io.Writer, format string, a ...interface{}) {
	Printf(PRINTF_FMT|PRINTF_LOG|FPRINTF_FMT, w, format, a...)
}
