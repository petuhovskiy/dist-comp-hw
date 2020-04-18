package importtool

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestCsv(t *testing.T) {
	f, err := os.Open("test_data.csv")
	assert.Nil(t, err)
	defer f.Close()

	r, err := NewCsvReader(f)
	assert.Nil(t, err)

	for {
		prod, err := r.Read()
		if err != nil {
			assert.Equal(t, io.EOF, err)
			break
		}

		spew.Dump(prod)
	}
}