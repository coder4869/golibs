// Copyright (c) 2016 coder4869 ( https://github.com/coder4869/golibs ). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*	gltest/glclient_test.go
 */

package test

import (
	"fmt"
	"testing"

	"github.com/coder4869/golibs/glnet"
)

/*	Testing case:
	usually the testing func's name makes up by 'Test+FunctionName';
	like TestGETReqWithURL(), 'GETReqWithURL' is testing func's name
*/
func TestGETReqWithURL(t *testing.T) {
	reqURL := "https://github.com/coder4869/golibs/blob/master/README.md"
	str, err := glnet.GETReqWithURL(reqURL)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(str)
	}
}

/*
	Parallel Test
*/
func Benchmark_GETReqWithURLParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			reqURL := "https://github.com/coder4869/golibs/blob/master/README.md"
			go glnet.GETReqWithURL(reqURL)
		}
	})
}

func Benchmark_GETReqWithURL(b *testing.B) {
	// must run b.N times. b.N will auto-adjust in running,
	// this ensure both time cost and caculated test data is reasonable.
	for i := 0; i < b.N; i++ {
		reqURL := "https://github.com/coder4869/golibs/blob/master/README.md"
		go glnet.GETReqWithURL(reqURL)
	}
}
