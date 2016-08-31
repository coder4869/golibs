// Copyright (c) 2016 coder4869 ( https://github.com/coder4869/golibs ). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*	glcrypto/glBase64.go
	provides base64 encoding and decoding operations of following results:
	[1].AES encrypto and decrypto
	[2].DES encrypto

Attention: the crypto related operations refer to “glcrypto/glAesDes.go”
*/

package glcrypto

import (
	"encoding/base64"

	"github.com/coder4869/golibs/glio"
)

/**********************************************************/
// AES-128 encrypt, then base64 encode. key length：16, 24, 32 bytes to AES-128, AES-192, AES-256
func Base64AesEnResult(orig []byte, key []byte) (string, error) {
	glio.FLPrintf("Before Aes Encrypt: %v\n", string(orig))
	aesresult, err := AesEncrypt([]byte(orig), key)
	if err != nil {
		//	panic(err)
		return "null", err
	}
	respDes := base64.StdEncoding.EncodeToString(aesresult)
	return respDes, nil
}

// base64 decode, then AES-128 decrypt. key length：16, 24, 32 bytes to AES-128, AES-192, AES-256
func Base64AesDeResult(crypted string, key []byte) (string, error) {
	glio.FLPrintf("Before Aes Decrypt: " + crypted)
	origAes, err := base64.StdEncoding.DecodeString(crypted)
	if err != nil {
		return "null", err
	}

	origStr, err := AesDecrypt(origAes, key)
	if err != nil {
		return "null", err
	}

	return string(origStr), nil
}

// base64 encoding result of DES
func Base64EncodeDesResult(orig []byte, key []byte) (string, error) {
	glio.FLPrintf("Before Des: " + string(orig))
	desresult, err := DesEncrypt([]byte(orig), key)
	if err != nil {
		return "null", err
	}
	respDes := base64.StdEncoding.EncodeToString(desresult)
	glio.FLPrintf(respDes)
	return respDes, nil
}
