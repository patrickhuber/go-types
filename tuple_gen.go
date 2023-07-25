//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"
)

var tupleTemplate string = `package types
{{range $index, $value := .}}
type Tuple{{$value}}[{{typeDef $value}} any] interface{
	tuple{{$value}}({{parameters $value}})
	Deconstruct() ({{typeDef $value}})
{{range $i := loop 1 $value -}}
{{print "\t"}}Value{{print $i}}() T{{print $i}}
{{end -}}
}

type tuple{{$value}}[{{typeDef $value}} any] struct{
{{range $i := loop 1 $value -}}
	{{print "\t"}}t{{$i}} T{{$i}}
{{end -}}
}

//lint:ignore U1000 method is used to implement Tuple{{$value}}[{{typeDef $value}}]
func (*tuple{{$value}}[{{typeDef $value}}]) tuple{{$value}}({{parameters $value}}) {}

func (t *tuple{{$value}}[{{typeDef $value}}]) Deconstruct() ({{typeDef $value}}){
	return {{range $index, $i := loop 1 $value}}{{if gt $index 0}}, {{end}}t.t{{$i}}{{end}}
}
{{range $i := loop 1 $value }}
func (t *tuple{{$value}}[{{typeDef $value}}]) Value{{$i}}() T{{$i}} {
	return t.t{{$i}}
}
{{end }}
func NewTuple{{$value}}[{{typeDef $value}} any]({{parameters $value}}) Tuple{{$value}}[{{typeDef $value}}]{
	return &tuple{{$value}}[{{typeDef $value}}]{
{{range $i := loop 1 $value -}}
		{{print "\t\t"}}t{{$i}}: t{{$i}}, 
{{end -}}
{{print "\t"}}}
}
{{end -}}
`

var tupleTestTesmplate = `package types_test
package types_test

import (
	"testing"

	"github.com/patrickhuber/go-types/tuple"
)

func TestTuple2(t *testing.T) {
	t.Run("deconstruct", func(t *testing.T) {
		expect1 := 1
		expect2 := "expect"
		t2 := tuple.New2(expect1, expect2)
		v1, v2 := t2.Deconstruct()
		if v1 != expect1 {
			t.Fatalf("expected %d found %d", expect1, v1)
		}
		if v2 != expect2 {
			t.Fatalf("expected %s found %s", expect2, v2)
		}
	})

	t.Run("value", func(t *testing.T) {
		expected1 := 1
		expected2 := 2
		tup := tuple.New2(expected1, expected2)
		if tup.Value1() != expected1 {
			t.Fatalf("expected value to equal %d", expected1)
		}
		if tup.Value2() != expected2 {
			t.Fatalf("expected value to equal %d", expected2)
		}
	})

}
`

func main() {
	funcs := template.FuncMap{
		"loop": func(low, high int) []int {
			var sequence []int
			for i := low; i <= high; i++ {
				sequence = append(sequence, i)
			}
			return sequence
		},
		"typeDef": func(value int) string {
			count := 0
			builder := &strings.Builder{}
			for i := 1; i <= value; i++ {
				if count > 0 {
					builder.WriteString(", ")
				}
				t := fmt.Sprintf("T%d", i)
				builder.WriteString(t)
				count++
			}
			return builder.String()
		},
		"parameters": func(value int) string {
			count := 0
			builder := &strings.Builder{}
			for i := 1; i <= value; i++ {
				if count > 0 {
					builder.WriteString(", ")
				}
				t := fmt.Sprintf("t%d T%d", i, i)
				builder.WriteString(t)
				count++
			}
			return builder.String()
		},
	}
	t, err := template.New("").Funcs(funcs).Parse(tupleTemplate)
	if err != nil {
		panic(err)
	}
	buf := &bytes.Buffer{}
	var values []int
	for i := 2; i <= 16; i++ {
		values = append(values, i)
	}
	err = t.Execute(buf, values)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("tuple.go", buf.Bytes(), 0666)
	if err != nil {
		panic(err)
	}
}
