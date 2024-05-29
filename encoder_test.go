// Copyright 2018 Alexander Poltoratskiy. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package fast_test

import (
	"bytes"
	"os"
	"reflect"
	"testing"

	"github.com/co11ter/goFAST"
)

var (
	encoder *fast.Encoder
	writer  *bytes.Buffer
)

func init() {
	ftpl, err := os.Open("testdata/test.xml")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = ftpl.Close()
	}()
	tpls, _ := fast.ParseXMLTemplate(ftpl)

	writer = &bytes.Buffer{}
	encoder = fast.NewEncoder(writer, tpls...)
	//encoder.SetLog(os.Stdout)
}

func encode(msg interface{}, expect []byte, t *testing.T) {
	err := encoder.Encode(msg)
	if err != nil {
		t.Fatal("can not encode", err)
	}

	if !reflect.DeepEqual(writer.Bytes(), expect) {
		t.Fatalf("data is not equal. current: %x expected: %x", writer.Bytes(), expect)
	}

	writer.Reset()
}

func TestDecimalEncode(t *testing.T) {
	encode(&decimalMessage1, decimalData1, t)
	// TODO have to provide decimal interface to set mantissa end exponent separately
	//encode(&decimalMessage2, decimalData2, t)
}

func TestSequenceEncode(t *testing.T) {
	encode(&sequenceMessage1, sequenceData1, t)
}

func TestByteVectorEncode(t *testing.T) {
	encode(&byteVectorMessage1, byteVectorData1, t)
}

func TestStringEncode(t *testing.T) {
	encode(&stringMessage1, stringData1, t)
}

func TestIntegerEncode(t *testing.T) {
	encode(&integerMessage1, integerData1, t)
}

func TestGroupEncode(t *testing.T) {
	encode(&groupMessage1, groupData1, t)
}

func TestReferenceEncode(t *testing.T) {
	encode(&referenceMessage1, referenceData1, t)
}

func TestOptionalTemplateIDEncode(t *testing.T) {
	encode(&testMessage1, testData1, t)
	encode(&testMessage2, testData2, t)
}
