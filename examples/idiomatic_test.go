package examples_test

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"testing"
)

func TestIdiomatic(t *testing.T) {
	tmpdir := t.TempDir()
	err := idiomatic(tmpdir)
	if err != nil {
		t.Fatal(err)
	}
}

func idiomatic(tmpdir string) error {
	d1 := []byte("hello\ngo\n")
	dat1 := path.Join(tmpdir, "dat1")

	err := os.WriteFile(dat1, d1, 0644)
	if err != nil {
		return err
	}

	dat2 := path.Join(tmpdir, "dat2")
	f, err := os.Create(dat2)
	if err != nil {
		return err
	}

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	if err != nil {
		return err
	}
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	if err != nil {
		return err
	}
	fmt.Printf("wrote %d bytes\n", n3)

	err = f.Sync()
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	if err != nil {
		return err
	}
	fmt.Printf("wrote %d bytes\n", n4)

	err = w.Flush()
	if err != nil {
		return err
	}

	return nil
}
