package examples

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"testing"

	"github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/handle"
	"github.com/patrickhuber/go-types/result"
)

// TestWrap shows file manipulation with a wrapped os package
func TestWrap(t *testing.T) {
	res := withWrap(t.TempDir())
	switch r := res.(type) {
	case types.Error[any]:
		t.Fatal(r.Error())
	case types.Ok[any]:
	}
}

func withWrap(directory string) (res types.Result[any]) {
	defer handle.Error(&res)
	d1 := []byte("hello\ngo\n")
	dat1 := path.Join(directory, "dat1")

	WriteFile(dat1, d1, 0644).Unwrap()

	dat2 := path.Join(directory, "dat2")
	f := Create(dat2).Unwrap()
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2 := Write(f, d2).Unwrap()
	fmt.Printf("wrote %d bytes\n", n2)

	n3 := WriteString(f, "writes\n").Unwrap()
	fmt.Printf("wrote %d bytes\n", n3)

	Sync(f).Unwrap()

	w := bufio.NewWriter(f)
	n4 := WriteString(w, "buffered\n").Unwrap()
	fmt.Printf("wrote %d bytes\n", n4)

	Flush(w).Unwrap()
	return
}

func WriteFile(name string, data []byte, perm fs.FileMode) types.Result[any] {
	return result.New[any](nil, os.WriteFile(name, data, perm))
}

func Create(name string) types.Result[*os.File] {
	return result.New(os.Create(name))
}

func Write(writer io.Writer, data []byte) types.Result[int] {
	return result.New(writer.Write(data))
}

func WriteString(writer io.StringWriter, data string) types.Result[int] {
	return result.New(writer.WriteString(data))
}

func Sync(f *os.File) types.Result[any] {
	return result.New[any](nil, f.Sync())
}

func Flush(w *bufio.Writer) types.Result[any] {
	return result.New[any](nil, w.Flush())
}
