// Copyright (c) 2021, Peter Ohler, All rights reserved.

package claujson

import (
	"claujson-go/claujson"
	"testing"

	"github.com/ohler55/ojg/oj"
)

func Test_ExampleTokenizer_Parse(t *testing.T) {
	toker := oj.Tokenizer{}
	h := claujson.Scanner{}
	src := `[true,null,123,12.3]{"x":12345678901234567890}`
	if err := toker.Parse([]byte(src), &h); err != nil {
		panic(err)
	}
	//fmt.Println(string(bytes.TrimSpace(h.buf)))

	// Output: [ true null 123 12.3 ] { x: 12345678901234567890 }
}
