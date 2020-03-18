package dialect

import (
	"reflect"
	"testing"
)

func TestDataTypeOf(t *testing.T) {
	dial := &sqlite3{}
	cases := []struct {
		Value  interface{}
		Expect string
	}{
		{"Tom", "text"},
		{123, "integer"},
		{1.2, "real"},
		{[]int{1, 2, 3}, "blob"},
	}

	for _, c := range cases {
		if typ := dial.DataTypeOf(reflect.ValueOf(c.Value)); typ != c.Expect {
			t.Fatalf("expect %s, but got %s", c.Expect, typ)
		}
	}
}
