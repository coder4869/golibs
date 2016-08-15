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
 */
