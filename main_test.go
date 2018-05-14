package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuine(t *testing.T) {
	old := os.Stdout // Backup of stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	outC := make(chan string)
	// Don't block
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// Return to normal state
	w.Close()
	os.Stdout = old
	out := <-outC

	// Read main.go
	f, err := ioutil.ReadFile("main.go")
	assert.NoError(t, err, "Failed to read main.go")

	// Assert output is the same as sourcecode
	assert.Equal(t, string(f), out)
}
