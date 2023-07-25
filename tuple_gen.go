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

var tupleTemplate string = `// this file was generated by "go generate" please do not edit
package types
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

var tupleTestTemplate = `// this file was generated by "go generate" please do not edit
package types_test

import (
	"testing"

	"github.com/patrickhuber/go-types"
)
{{range $index, $value := .}}
func TestTuple{{$value}}(t *testing.T) {
	t.Run("deconstruct", func(t *testing.T) {
		{{- range $i, $v := loop 1 $value}}
		expect{{$v}} := {{$i}}
		{{- end}}
		t{{$value}} := types.NewTuple{{$value}}({{range $i, $v := loop 1 $value}}{{if gt $i 0}}, {{end}}expect{{$v}}{{end}})
		{{range $i, $v := loop 1 $value}}{{if gt $i 0}}, {{end}}v{{$v}}{{end}} := t{{$value}}.Deconstruct()
		{{range $i, $v := loop 1 $value}}
		if v{{$v}} != expect{{$v}} {
			t.Fatalf("expected %d found %d", expect{{$v}}, v{{$v}})
		}{{end}}		
	})

	t.Run("value", func(t *testing.T) {
		{{- range $i, $v := loop 1 $value}}
		expect{{$v}} := {{$i}}
		{{- end}}
		tup := types.NewTuple{{$value}}({{range $i, $v := loop 1 $value}}{{if gt $i 0}}, {{end}}expect{{$v}}{{end}})
		{{range $i, $v := loop 1 $value}}
		if tup.Value{{$v}}() != expect{{$v}} {
			t.Fatalf("expected %d to equal %d", tup.Value{{$v}}(), expect{{$v}})
		}{{end}}	
	})
}
{{end -}}
`

var tupleFuncsTemplate = `package tuple

import "github.com/patrickhuber/go-types"
{{range $index, $value := .}}
func New{{$value}}[{{typeDef $value}} any]({{parameters $value}}) types.Tuple{{$value}}[{{typeDef $value}}] {
	return types.NewTuple{{$value}}({{range $i, $v := loop 1 $value}}{{if gt $i 0}}, {{end}}t{{$v}}{{end}})
}
{{end}}
`

var tupleFuncsTestTemplate = `package tuple_test

import (
	"github.com/patrickhuber/go-types/tuple"

	"testing"
)
{{range $index, $value := .}}
func TestNew{{$value}}(t *testing.T) {
	t.Run("value", func(t *testing.T) {
		{{- range $i, $v := loop 1 $value}}
		expect{{$v}} := {{$i}}
		{{- end}}
		tup := tuple.New{{$value}}({{range $i, $v := loop 1 $value}}{{if gt $i 0}}, {{end}}expect{{$v}}{{end}})
		{{range $i, $v := loop 1 $value}}
		if tup.Value{{$v}}() != expect{{$v}} {
			t.Fatalf("expected %d to equal %d", tup.Value{{$v}}(), expect{{$v}})
		}{{end}}	
	})
}
{{end}}
`

var resultFuncsTemplate = `package result

import (
	"github.com/patrickhuber/go-types"
)

func New[T any](ok T, err error) types.Result[T] {
	return types.NewResult(ok, err)
}
{{range $index, $value := .}}
func New{{$value}}[{{typeDef $value}} any]({{parameters $value}}, err error) types.Result[types.Tuple{{$value}}[{{typeDef $value}}]] {
	return types.NewResult(
		types.NewTuple{{$value}}({{values $value}}), 
		err,
	)
}
{{end}}`

var resultFuncsTestTemplate = `package result_test

import (
	"testing"

	"github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/result"
)

func TestNew(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		switch result.New(1, nil).(type) {
		case types.Error[int]:
			t.Fatalf("expected types.Ok[int]")
		}
	})
	t.Run("func", func(t *testing.T) {
		f := func() (int, error) {
			return 1, nil
		}
		switch result.New(f()).(type) {
		case types.Error[int]:
			t.Fatalf("expected types.Ok[int]")
		}
	})
}
{{range $index, $value := .}}
func TestNew{{$value}}(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		switch result.New{{$value}}({{range $i, $v := loop 1 $value }}{{if $i}}, {{end}}{{$v}}{{end}}, nil).(type) {
		case types.Error[types.Tuple{{$value}}[{{join ", " (repeat "int" $value)}}]]:
			t.Fatalf("expected types.Ok[types.Tuple{{$value}}[{{join ", " (repeat "int" $value)}}]]")
		}
	})
	t.Run("func", func(t *testing.T) {
		f := func() ({{join ", " (repeat "int" $value)}}, error) {
			return {{range $i, $v := loop 1 $value}}{{if $i}}, {{end}}{{$v}}{{end}}, nil
		}
		switch result.New{{$value}}(f()).(type) {
		case types.Error[types.Tuple{{$value}}[{{join ", " (repeat "int" $value)}}]]:
			t.Fatalf("expected types.Ok[Tuple{{$value}}[{{join ", " (repeat "int" $value)}}]]")
		}
	})
}
{{end}}`

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
		"values": func(value int) string {
			count := 0
			builder := &strings.Builder{}
			for i := 1; i <= value; i++ {
				if count > 0 {
					builder.WriteString(", ")
				}
				t := fmt.Sprintf("t%d", i)
				builder.WriteString(t)
				count++
			}
			return builder.String()
		},
		"join": func(sep string, others []string) string {
			return strings.Join(others, sep)
		},
		"repeat": func(s string, count int) []string {
			var strList []string
			for i := 0; i < count; i++ {
				strList = append(strList, s)
			}
			return strList
		},
	}

	m := map[string]string{
		tupleTemplate:           "tuple.go",
		tupleTestTemplate:       "tuple_test.go",
		tupleFuncsTemplate:      "tuple/new.go",
		tupleFuncsTestTemplate:  "tuple/new_test.go",
		resultFuncsTemplate:     "result/new.go",
		resultFuncsTestTemplate: "result/new_test.go",
	}
	for tmpl, outFile := range m {
		err := Generate(funcs, tmpl, outFile)
		handle(err)
	}
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func Generate(funcs template.FuncMap, tmpl string, outputFile string) error {
	t, err := template.New("").Funcs(funcs).Parse(tmpl)
	if err != nil {
		return err
	}
	buf := &bytes.Buffer{}
	var values []int
	for i := 2; i <= 16; i++ {
		values = append(values, i)
	}
	err = t.Execute(buf, values)
	if err != nil {
		return err
	}
	err = os.WriteFile(outputFile, buf.Bytes(), 0666)
	return err
}
