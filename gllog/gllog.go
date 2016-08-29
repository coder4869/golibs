// Copyright (c) 2016 coder4869 ( https://github.com/coder4869/golibs ). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*	gllog/gllog.go
 */

package gllog

import (
	"fmt"
	"log"
	"os"
	//	"strings"
	//	"path"
	"runtime"
	"time"
)

var (
	LogFile          *os.File = nil
	logPrintFileName string
	logPrintFileLine int
)

/*
perm means unix permission bits, which can be a number of 4 characters, like 0777.
each character makes up by 4(r-read), 2(w-write) and 1(x-excute). for example 7=4+2+1.
0777 equals -rwxrwxrwx, meaning as following permissions to current file(or dir):
first 7 means current user has read(4)+write(2)+excute(1) permissions,
second 7 means user's group has read(4)+write(2)+excute(1) permissions,
last 7 means other user has read(4)+write(2)+excute(1) permissions.
*/
func NewLog() (*os.File, bool) {
	//create log dir
	logDirCreateErr := os.MkdirAll("./logs/serv", 0744)
	if logDirCreateErr != nil {
		log.Fatal("Create log dir failed!\n ERROR_INFO: ", logDirCreateErr)
		panic(logDirCreateErr)
		return nil, false
	}

	//create log file
	t := time.Now()
	//	filepath := "./logs/serv/log_" + strings.Replace(t.String()[:19], ":", "_", 3) + ".log"
	filepath := "./logs/serv/log_" + t.String() + ".log"
	file, openFileErr := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if openFileErr != nil {
		log.Fatal("Open log file failed!\n ERROR_INFO: ", openFileErr)
		panic(openFileErr)
		return nil, false
	}
	defer file.Close()

	return file, true
}

//appending logs to file "LogFile"
func AppendLog(logs string) {

	if LogFile == nil { //log file not set
		newLog, succeed := NewLog() //create new log file
		if !succeed {
			return
		}
		LogFile = newLog
	} else {
		logFileInfo, logExistErr := os.Stat(LogFile.Name())
		if os.IsNotExist(logExistErr) { //file not exist
			newLog, succeed := NewLog() //create new log file
			if !succeed {
				return
			}
			LogFile = newLog
		} else { //if size of log file is too big(over 20MB), create new log file
			if logFileInfo.Size() > 20*1024*1024 {
				newLog, succeed := NewLog() //create new log file
				if !succeed {
					return
				}
				LogFile = newLog
			}
		}
	}

	logFile, logOpenErr := os.OpenFile(LogFile.Name(), os.O_RDWR|os.O_APPEND, 0644)
	if logOpenErr != nil {
		fmt.Printf("Open log file error=%s\r\n", logOpenErr.Error())
		os.Exit(-1)
	}
	defer logFile.Close()

	logger := log.New(logFile, "", log.Ldate|log.Ltime)
	logger.Printf(logs)
}

func SetCurrentFileName() {
	_, logPrintFileName, logPrintFileLine, _ := runtime.Caller(0)
	fmt.Printf("logPrintFileName=%v,logPrintFileLine=%d\n", logPrintFileName, logPrintFileLine)
}
