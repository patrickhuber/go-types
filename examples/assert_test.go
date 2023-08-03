package examples_test

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/assert"
	"github.com/patrickhuber/go-types/handle"
)

func TestAssert(t *testing.T) {
	res := withAssert(t.TempDir())
	switch r := res.(type) {
	case types.Error[any]:
		t.Fatal(r.Value)
	case types.Ok[any]:
	}
}

func withAssert(directory string) (res types.Result[any]) {

	// when assert panics,
	// the handler will recover and set the error result
	defer handle.Error(&res)

	d1 := []byte("hello\ngo\n")
	dat1 := path.Join(directory, "dat1")

	err := os.WriteFile(dat1, d1, 0644)
	assert.Nil(err)

	dat2 := path.Join(directory, "dat2")
	f, err := os.Create(dat2)
	assert.Nil(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	assert.Nil(err)

	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	assert.Nil(err)
	fmt.Printf("wrote %d bytes\n", n3)

	err = f.Sync()
	assert.Nil(err)

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	assert.Nil(err)
	fmt.Printf("wrote %d bytes\n", n4)

	err = w.Flush()
	assert.Nil(err)

	return
}
